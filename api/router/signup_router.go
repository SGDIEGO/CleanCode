package router

import (
	"github.com/SGDIEGO/CleanCode/api/handler"
	"github.com/SGDIEGO/CleanCode/internal/port"
	"github.com/SGDIEGO/CleanCode/internal/service"
	"github.com/gin-gonic/gin"
)

type SignUpRouter struct {
	GroupRoute *gin.RouterGroup
	URepo      port.UserRepo
}

func NewSignUpR(GroupRoute *gin.RouterGroup, URepo port.UserRepo) port.RouterI {
	return &SignUpRouter{
		GroupRoute: GroupRoute,
		URepo:      URepo,
	}
}

func (s *SignUpRouter) Load(route string) {
	service := service.NewsignupService(s.URepo)
	sHl := handler.NewSignupHandler(service)

	router := s.GroupRoute.Group(route)
	{
		router.GET("/", sHl.Index)
		router.POST("/", sHl.SignUp)
	}
}
