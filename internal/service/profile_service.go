package service

import (
	"strconv"
	"time"

	"github.com/SGDIEGO/CleanCode/internal/domain/entity"
	"github.com/SGDIEGO/CleanCode/internal/port"
	"github.com/SGDIEGO/CleanCode/pkg/util"
	"github.com/gin-gonic/gin"
)

type ProfileService interface {
	Principal(c *gin.Context, secretKey []byte) (*entity.User, error)
	GetUser(c *gin.Context, secretKey []byte) (*entity.User, error)
	GroupsProfile(c *gin.Context, secretKey []byte) (*[]entity.Group, error)
	CreateGroup(groupNew *entity.Group, userAdmin *entity.User) error
	AddUserToGroup(assignInfo *entity.Assignment) error
	UsersFromGroup(groupId string) (*[]entity.User, error)
	DeleteUserFromGroup(userId, groupId int) error
}

type profileService struct {
	UserRepo   port.UserRepo
	GroupRepo  port.GroupRepo
	AssignRepo port.AssignmentRepo
}

func NewprofileService(UserRepo port.UserRepo, GroupRepo port.GroupRepo, AssignRepo port.AssignmentRepo) ProfileService {
	return &profileService{
		UserRepo:   UserRepo,
		GroupRepo:  GroupRepo,
		AssignRepo: AssignRepo,
	}
}

func (p *profileService) GetUser(c *gin.Context, secretKey []byte) (*entity.User, error) {

	request := util.ExtractToken(c)

	return util.GetTokenValue(request, secretKey)
}

func (p *profileService) Principal(c *gin.Context, secretKey []byte) (*entity.User, error) {

	return p.GetUser(c, secretKey)
}

func (p *profileService) GroupsProfile(c *gin.Context, secretKey []byte) (*[]entity.Group, error) {

	user, err := p.GetUser(c, secretKey)

	if err != nil {
		return nil, err
	}

	return p.GroupRepo.GetGroupsByAdminName(user.UserName)
}

func (p *profileService) CreateGroup(groupNew *entity.Group, userAdmin *entity.User) error {

	groupNew.Admin = userAdmin.UserName
	groupNew.Created = time.Now().Format("2006-01-02 15:04:05")
	groupNew.Members = 1

	return p.GroupRepo.CreateGroup(groupNew)
}

func (p *profileService) AddUserToGroup(assignInfo *entity.Assignment) error {

	// Get user with id
	user, err := p.UserRepo.GetUserById(assignInfo.UserId)

	// Verify if user or an err exist
	if user == nil {
		return err
	}

	// Get group with id
	group, err := p.GroupRepo.GetGroupById(assignInfo.GroupId)
	// Verify if group or an err exist
	if group == nil {
		return err
	}

	// Else assing user to group
	return p.AssignRepo.UserToGroup(assignInfo)
}

func (p *profileService) userById(id int, c chan entity.User) {
	user, _ := p.UserRepo.GetUserById(id)
	c <- *user
}

func (p *profileService) UsersFromGroup(groupId string) (*[]entity.User, error) {

	id, err := strconv.Atoi(groupId)
	if err != nil {
		return nil, err
	}

	userIds, err := p.GroupRepo.UserIdsFromGroup(id)
	if err != nil {
		return nil, err
	}

	var Users []entity.User
	UserChan := make(chan entity.User)

	for _, id := range userIds {
		// Concurrency
		go p.userById(id, UserChan)
		// Get value from chanel
		user := <-UserChan
		// Add value to list
		Users = append(Users, user)
	}

	close(UserChan)

	return &Users, nil
}

func (p *profileService) DeleteUserFromGroup(userId, groupId int) error {

	// Verify if user exists
	user, err := p.UserRepo.GetUserById(userId)
	if (user == nil) || err != nil {
		return err
	}

	return p.AssignRepo.DeleteAssignment(userId, groupId)
}
