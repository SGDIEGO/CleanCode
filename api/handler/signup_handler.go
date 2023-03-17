package handler

import (
	"net/http"

	"github.com/SGDIEGO/CleanCode/internal/domain/entity"
	"github.com/SGDIEGO/CleanCode/internal/service"
	"github.com/gin-gonic/gin"
)

type SignupHandler struct {
	service service.SignupService
}

func NewSignupHandler(service service.SignupService) *SignupHandler {
	return &SignupHandler{
		service: service,
	}
}

func (s *SignupHandler) Index(c *gin.Context) {

}

func (s *SignupHandler) SignUp(c *gin.Context) {
	var UserSign entity.User

	// If user is bad registered
	if err := c.ShouldBindJSON(&UserSign); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Signing user
	if err := s.service.SignUp(&UserSign); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"response": "User registerd succsesfully",
	})

}
