package server

import (
	"net/http"

	"github.com/Superm4n97/account-server/pkg/account/permission"
	"github.com/gin-gonic/gin"
	"k8s.io/klog/v2"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", helloWorld)
	v1 := router.Group("/api/v1")
	GroupPermission := router.Group("/api/v1/permission")
	permission.UpdatePermission(GroupPermission.Group("/"))
	userRouter(v1.Group("/user"))
	return router
}

func helloWorld(ctx *gin.Context) {
	klog.V(2)
	klog.Infof("the server is running and it's helloing you!!!!")
	ctx.IndentedJSON(http.StatusOK, "Hello World from server")
}
