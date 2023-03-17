package router

import (
	"github.com/SGDIEGO/CleanCode/api/handler"
	"github.com/SGDIEGO/CleanCode/internal/core/config"
	"github.com/SGDIEGO/CleanCode/internal/port"
	"github.com/SGDIEGO/CleanCode/internal/service"
	"github.com/gin-gonic/gin"
)

type LoginRouter struct {
	GroupRoute *gin.RouterGroup
	configT    *config.JWTConfig
	URepo      port.UserRepo
}

func NewLoginR(GroupRoute *gin.RouterGroup, configT *config.JWTConfig, URepo port.UserRepo) port.RouterI {
	return &LoginRouter{
		GroupRoute: GroupRoute,
		configT:    configT,
		URepo:      URepo,
	}
}

func (r *LoginRouter) Load(route string) {

	/*DEPENDENDIES*/
	// Service dependency (inyect db)
	service := service.NewloginService(r.URepo)
	// Get controller (inyect service)
	lHl := handler.NewLoginHandler(service, r.configT)

	router := r.GroupRoute.Group(route)
	{
		router.GET("/", lHl.Index)
		router.POST("/", lHl.Login)
	}
}
