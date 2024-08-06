package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type server struct {
	router *gin.Engine
}

func newServer() {
	router := gin.Default()
	router.GET("/hello-world", pong)
}

func pong(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "pong from server")
}
