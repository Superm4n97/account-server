package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"k8s.io/klog/v2"
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
