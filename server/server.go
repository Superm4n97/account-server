package server

import (
	"github.com/Superm4n97/account-server/pkg/database/mongodb"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Router       *gin.Engine
	Port         string
	DatabaseName string
	DatabaseURI  string
}

func (s *Server) Start() error {
	if err := mongodb.Init(s.DatabaseURI, s.DatabaseName); err != nil {
		return err
	}

	if err := s.Router.Run(s.Port); err != nil {
		return err
	}
	return nil
}
