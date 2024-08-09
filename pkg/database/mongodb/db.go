package mongodb

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"k8s.io/klog/v2"
	"strings"
)

const (
	defaultName = "account-server"
	defaultUri  = "mongodb://localhost:27017"
)

var (
	database *mongo.Database
)

// dbUri specifies the database connection string and dbName specifies the
// default values:
//
//	uri: mongodb://localhost:27017
//	dbName: account-server
func Init(dbUri, dbName string) error {
	if dbUri == "" {
		dbUri = defaultUri
	}
	if dbName == "" {
		dbName = defaultName
	}

	clientOpts := options.Client().ApplyURI(dbUri)

	var err error
	c, err := mongo.Connect(context.Background(), clientOpts)
	if err != nil {
		klog.Error("failed to apply create database client")
		return err
	}
	database = c.Database(dbName)
	return nil
}

func Add(document any, collection string) (string, error) {
	res, err := database.Collection(collection).InsertOne(context.Background(), document)
	if err != nil {
		klog.Errorf("failed to add document %v to mongodb collection %v\n", document, collection)
		return "", err
	}
	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func Get(filters map[string]string, collection string) ([]byte, error) {
	var cur *mongo.Cursor
	var err error

	bsonFilter := make(bson.D, 0)
	for k, v := range filters {
		bsonFilter = append(bsonFilter, bson.E{
			Key:   strings.ToLower(k),
			Value: v,
		})
	}

	cur, err = database.Collection(collection).Find(context.Background(), bsonFilter)

	var results []bson.M
	for cur.Next(context.Background()) {
		var res bson.M
		if err = cur.Decode(&res); err != nil {
			return nil, err
		}
		results = append(results, res)
	}

	return json.MarshalIndent(results, "", "    ")
}

func Delete(filters map[string]string, collection string) error {
	bsonFilter := make(bson.D, 0)
	for k, v := range filters {
		bsonFilter = append(bsonFilter, bson.E{
			Key:   strings.ToLower(k),
			Value: v,
		})
	}

	res := database.Collection(collection).FindOneAndDelete(context.Background(), bsonFilter)
	return res.Err()
}

// IfPresent returns
// nil,false,err => if any error occur
// nil,false,nil => if there is no object found with current filter
// []byte{},true,nil => if there is one or many objects found
func IfPresent(filter map[string]string, collection string) ([]byte, bool, error) {
	data, err := Get(filter, collection)
	if err != nil {
		return nil, false, err
	}

	if string(data) == "null" {
		return nil, false, nil
	}
	return data, true, nil
}
