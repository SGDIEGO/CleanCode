package router

import (
	"database/sql"

	"github.com/SGDIEGO/CleanCode/api/middleware"
	"github.com/SGDIEGO/CleanCode/internal/core/config"
	"github.com/SGDIEGO/CleanCode/internal/domain/repository"
	"github.com/gin-gonic/gin"
)

const ROOTH_PATH = ""

func Setup(config *config.SvConfig, server *gin.Engine, db *sql.DB) {

	publicRouter := server.Group(ROOTH_PATH)
	privateRouter := server.Group(ROOTH_PATH)
	privateRouter.Use(middleware.JwtAuthMiddleware(config.JWT.Key))

	// Repositories
	var (
		UserRepository   = repository.NewuserRepo(db)
		GroupRepository  = repository.NewGroupRepo(db)
		AssingRepository = repository.NewAssignmentRepo(db)
	)

	// Routers
	// Public Routes
	var (
		HomeR   = NewHomeR(publicRouter, UserRepository)
		LoginR  = NewLoginR(publicRouter, &config.JWT, UserRepository)
		SignupR = NewSignUpR(publicRouter, UserRepository)
	)

	// Private Routes
	var (
		//Add JWT Auth middleware
		ProfileR = NewProfileR(privateRouter, &config.JWT, GroupRepository, UserRepository, AssingRepository)
	)

	// Load Routers
	HomeR.Load("/home")
	LoginR.Load("/login")
	ProfileR.Load("/profile")
	SignupR.Load("/signup")
}
