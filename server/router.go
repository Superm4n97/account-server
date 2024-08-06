package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/hello-world", helloWorld)
	v1 := router.Group("/api/v1")
	userRouter(v1.Group("/user"))

	return router
}

func helloWorld(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, "Hello World from server")
}
