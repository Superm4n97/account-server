package account

import (
	"github.com/Superm4n97/account-server/pkg/database/mongodb"
	"github.com/gin-gonic/gin"
	"k8s.io/klog/v2"
	"net/http"
)

const (
	collectionUser = "user"
)

func CreateUser(ctx *gin.Context) {
	var u User
	if err := ctx.BindJSON(&u); err != nil {
		klog.Error(err.Error())
		ctx.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	uid, err := mongodb.Add(u, collectionUser)
	if err != nil {
		klog.Error(err.Error())
		ctx.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	klog.Infof("user successfully inserted with userid %s", uid)
	ctx.IndentedJSON(http.StatusOK, "user added successfully")
}
