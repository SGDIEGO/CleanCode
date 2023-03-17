package service

import (
	"errors"

	"github.com/SGDIEGO/CleanCode/internal/domain/entity"
	"github.com/SGDIEGO/CleanCode/internal/port"
	"github.com/SGDIEGO/CleanCode/pkg/util"
)

type LoginService interface {
	Login(user *entity.User) (*entity.User, error)
	CreateToken(UserLog *entity.User, secret []byte) (string, error)
}

type loginService struct {
	UserRepo port.UserRepo
}

func NewloginService(UserRepo port.UserRepo) LoginService {
	return &loginService{
		UserRepo: UserRepo,
	}
}

func (u *loginService) Login(user *entity.User) (*entity.User, error) {
	exists, err := u.UserRepo.GetUserByEmail(user.Email)

	if (err != nil) || (exists == nil) {
		return nil, err
	}

	if exists.Password != user.Password {
		return nil, errors.New("password is incorrect")
	}

	return exists, nil
}

func (u *loginService) CreateToken(UserLog *entity.User, secret []byte) (string, error) {

	return util.CreateToken(UserLog, secret)
}
