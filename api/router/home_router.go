package router

import (
	"github.com/SGDIEGO/CleanCode/api/handler"
	"github.com/SGDIEGO/CleanCode/internal/port"
	"github.com/SGDIEGO/CleanCode/internal/service"
	"github.com/gin-gonic/gin"
)

type HomeRouter struct {
	GroupRoute *gin.RouterGroup
	URepo      port.UserRepo
}

func NewHomeR(GroupRoute *gin.RouterGroup, URepo port.UserRepo) port.RouterI {
	return &HomeRouter{
		GroupRoute: GroupRoute,
		URepo:      URepo,
	}
}

func (r *HomeRouter) Load(route string) {
	service := service.NewHomeService(r.URepo)
	hHl := handler.NewHomeHandler(service)

	router := r.GroupRoute.Group(route)
	{
		router.GET("/", hHl.Index)
	}
}
