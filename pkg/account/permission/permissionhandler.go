package permission

import (
	"encoding/json"
	"net/http"

	"github.com/Superm4n97/account-server/pkg/account/user"
	"github.com/Superm4n97/account-server/pkg/database/mongodb"
	"github.com/gin-gonic/gin"
	"k8s.io/klog/v2"
)

func UpdateRole(ctx *gin.Context) {
	var CurrentPermissionRef PermissionRef
	ctx.BindJSON(&CurrentPermissionRef)
	filter := map[string]string{}
	filter[CurrentPermissionRef.UserId] = CurrentPermissionRef.UserId
	data, err := mongodb.Get(filter, "user")
	if err != nil {
		klog.Error(err.Error())
		ctx.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	var CurrentUser []user.User
	if err = json.Unmarshal(data, &CurrentUser); err != nil {
		klog.Error(err.Error())
		ctx.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	//variable CurrentUser array has only one user
	CurrentUser[0].Permissions[CurrentPermissionRef.ClusterID] = CurrentPermissionRef.Role

	ctx.IndentedJSON(http.StatusOK, "The Role has been added to the user")
}
