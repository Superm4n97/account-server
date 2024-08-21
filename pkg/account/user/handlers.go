package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Superm4n97/account-server/pkg/database/mongodb"
	"github.com/Superm4n97/account-server/pkg/util"
	"github.com/gin-gonic/gin"
	"k8s.io/klog/v2"
)

const (
	collectionUser = "user"
	keyId          = "id"
	keyEmail       = "email"
	keyName        = "name"
)

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
	//Permissions is a map of map[ClusterID]Role
	Permissions map[string]string
}

func CreateUser(ctx *gin.Context) {
	var u User
	if err := ctx.BindJSON(&u); err != nil {
		klog.Error(err.Error())
		ctx.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	_, found, err := mongodb.IfPresent(map[string]string{
		keyEmail: u.Email,
	}, collectionUser)
	if err != nil {
		klog.Errorf(err.Error())
		ctx.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	if found {
		err = fmt.Errorf("email: %s is used for another account", u.Email)
		klog.Errorf(err.Error())
		ctx.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	u.ID, err = util.GetUniqueID()
	if err != nil {
		klog.Error(err.Error())
		ctx.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	_, err = mongodb.Add(u, collectionUser)
	if err != nil {
		klog.Error(err.Error())
		ctx.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	klog.Infof("user successfully inserted with userid %s", u.ID)
	ctx.IndentedJSON(http.StatusOK, "user added successfully")
}

func GetUsers(ctx *gin.Context) {
	filter := map[string]string{}
	if userId, found := ctx.Params.Get("user"); found {
		filter[keyId] = userId
	}

	data, err := mongodb.Get(filter, collectionUser)
	if err != nil {
		klog.Error(err.Error())
		ctx.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	var usrs []User
	if err = json.Unmarshal(data, &usrs); err != nil {
		klog.Error(err.Error())
		ctx.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	klog.Info("user listed successfully")
	ctx.IndentedJSON(http.StatusOK, usrs)
}

func DeleteUser(ctx *gin.Context) {
	filter := map[string]string{}
	if userId, exists := ctx.Params.Get(":user"); exists {
		filter[keyId] = userId
	}
	if err := mongodb.Delete(filter, collectionUser); err != nil {
		klog.Error(err.Error())
		ctx.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	msg := "user deleted successfully"
	klog.Info(msg)
	ctx.IndentedJSON(http.StatusOK, msg)
}
