package router

import (
	"github.com/SGDIEGO/CleanCode/api/handler"
	"github.com/SGDIEGO/CleanCode/internal/core/config"
	"github.com/SGDIEGO/CleanCode/internal/port"
	"github.com/SGDIEGO/CleanCode/internal/service"
	"github.com/gin-gonic/gin"
)

type ProfileRouter struct {
	GroupRoute *gin.RouterGroup
	config     *config.JWTConfig
	GRepo      port.GroupRepo
	URepo      port.UserRepo
	ARepo      port.AssignmentRepo
}

func NewProfileR(GroupRoute *gin.RouterGroup, config *config.JWTConfig, GRepo port.GroupRepo, URepo port.UserRepo, ARepo port.AssignmentRepo) port.RouterI {
	return &ProfileRouter{
		GroupRoute: GroupRoute,
		config:     config,
		GRepo:      GRepo,
		ARepo:      ARepo,
		URepo:      URepo,
	}
}

func (r *ProfileRouter) Load(route string) {

	/*DEPENDENDIES*/
	// Service dependency (inyect repos)
	service := service.NewprofileService(r.URepo, r.GRepo, r.ARepo)
	// Get controller (inyect service)
	pHl := handler.NewProfileHandler(service, r.config)

	router := r.GroupRoute.Group(route)
	{
		router.GET("/", pHl.Principal)
		router.GET("/groups", pHl.GroupsProfile)
		router.POST("/groups", pHl.CreateGroup)
		router.GET("/groups/:id", pHl.UsersFromGroup)
		router.POST("/groups/:id", pHl.AddUserToGroup)
		router.DELETE("/groups/:id", pHl.DeleteUserFromGroup)
	}
}
