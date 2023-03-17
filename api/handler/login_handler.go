package handler

import (
	"net/http"

	"github.com/SGDIEGO/CleanCode/internal/core/config"
	"github.com/SGDIEGO/CleanCode/internal/domain/entity"
	"github.com/SGDIEGO/CleanCode/internal/service"
	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	service service.LoginService
	configT *config.JWTConfig
}

func NewLoginHandler(service service.LoginService, configT *config.JWTConfig) *LoginHandler {
	return &LoginHandler{
		service: service,
		configT: configT,
	}
}

func (p *LoginHandler) Index(c *gin.Context) {

}

func (p *LoginHandler) Login(c *gin.Context) {

	var UserLog entity.User

	// If user is bad
	if err := c.ShouldBindJSON(&UserLog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error1": err.Error(),
		})
		return
	}

	// If user was not registered
	UserInfo, err := p.service.Login(&UserLog)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error2": err.Error(),
		})
		return
	}

	// If registered
	AuthCookie, err := p.service.CreateToken(UserInfo, p.configT.Key)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error3": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"AuthCookie": AuthCookie,
	})
}
