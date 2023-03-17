package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/SGDIEGO/CleanCode/internal/core/config"
	"github.com/SGDIEGO/CleanCode/internal/domain/entity"
	"github.com/SGDIEGO/CleanCode/internal/service"
	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
	service service.ProfileService
	configT *config.JWTConfig
}

func NewProfileHandler(service service.ProfileService, configT *config.JWTConfig) *ProfileHandler {
	return &ProfileHandler{
		service: service,
		configT: configT,
	}
}

func (p *ProfileHandler) Principal(c *gin.Context) {

	user, err := p.service.Principal(c, p.configT.Key)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errorHeader": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func (p *ProfileHandler) GroupsProfile(c *gin.Context) {

	groups, err := p.service.GroupsProfile(c, p.configT.Key)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"groups": groups,
	})
}

func (p *ProfileHandler) CreateGroup(c *gin.Context) {

	var GroupN entity.Group

	if err := c.ShouldBindJSON(&GroupN); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := p.service.GetUser(c, p.configT.Key)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := p.service.CreateGroup(&GroupN, user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"response": "group created",
	})
}

func (p *ProfileHandler) AddUserToGroup(c *gin.Context) {

	// Get id from the group
	idGroup := c.Param("id")

	// Get the user info
	var AssingInfo entity.Assignment
	if err := c.ShouldBindJSON(&AssingInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Add user to group
	err := p.service.AddUserToGroup(&AssingInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": fmt.Sprintln("user", AssingInfo.UserId, "added to group", idGroup),
	})
}

func (p *ProfileHandler) UsersFromGroup(c *gin.Context) {
	// Get id param
	groupId := c.Param("id")
	// Use service to add user with this id
	users, err := p.service.UsersFromGroup(groupId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (p *ProfileHandler) DeleteUserFromGroup(c *gin.Context) {
	// Get id param
	idgroup, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// Bind user id data
	userId, err := strconv.Atoi(c.Request.FormValue("id"))
	if err != nil {
		c.JSON(http.StatusNotAcceptable, err.Error())
		return
	}

	// Delete user in group id
	if err := p.service.DeleteUserFromGroup(userId, idgroup); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusBadRequest, string("User deleted"))
}
