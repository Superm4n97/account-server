package server

import (
	"github.com/Superm4n97/account-server/pkg/account"
	"github.com/gin-gonic/gin"
)

func userRouter(r *gin.RouterGroup) {
	r.POST("/", account.CreateUser)
}
