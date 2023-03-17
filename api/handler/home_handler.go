package handler

import (
	"net/http"

	"github.com/SGDIEGO/CleanCode/internal/service"
	"github.com/gin-gonic/gin"
)

type HomeHandler struct {
	service service.HomeService
}

func NewHomeHandler(service service.HomeService) *HomeHandler {
	return &HomeHandler{
		service: service,
	}
}

func (p *HomeHandler) Index(c *gin.Context) {

	users, err := p.service.Index()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})

}
