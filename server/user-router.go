package server

import (
	"github.com/Superm4n97/account-server/pkg/account/user"
	"github.com/gin-gonic/gin"
)

func userRouter(r *gin.RouterGroup) {
	r.POST("/", user.CreateUser)
	r.GET("/", user.GetUsers)
	r.GET("/:user", user.GetUsers)
	r.DELETE("/:user", user.DeleteUser)
}
