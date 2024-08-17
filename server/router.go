package server

import (
	"github.com/gin-gonic/gin"
	"k8s.io/klog/v2"
	"net/http"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", helloWorld)
	v1 := router.Group("/api/v1")
	userRouter(v1.Group("/user"))

	return router
}

func helloWorld(ctx *gin.Context) {
	klog.V(2)
	klog.Infof("the server is running and it's helloing you!!!!")
	ctx.IndentedJSON(http.StatusOK, "Hello World from server")
}
