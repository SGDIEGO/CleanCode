package app

import (
	"database/sql"
	"fmt"

	"github.com/SGDIEGO/CleanCode/api/router"
	"github.com/SGDIEGO/CleanCode/internal/core/config"
	"github.com/gin-gonic/gin"
)

type Server struct {
	SvConfig  *config.SvConfig
	GinServer *gin.Engine
}

func NewServer(SvConfig *config.SvConfig) *Server {
	return &Server{
		SvConfig:  SvConfig,
		GinServer: gin.Default(),
	}
}

func (s *Server) Load(db *sql.DB) {

	// Load routers
	router.Setup(s.SvConfig, s.GinServer, db)

	// Run server
	s.GinServer.Run(fmt.Sprint(":", s.SvConfig.Port))
}
