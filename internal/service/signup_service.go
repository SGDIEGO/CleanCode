package service

import (
	"github.com/SGDIEGO/CleanCode/internal/domain/entity"
	"github.com/SGDIEGO/CleanCode/internal/port"
)

type SignupService interface {
	SignUp(user *entity.User) error
}

type signupService struct {
	UserRepo port.UserRepo
}

func NewsignupService(UserRepo port.UserRepo) SignupService {
	return &signupService{
		UserRepo: UserRepo,
	}
}

func (s *signupService) SignUp(user *entity.User) error {

	return s.UserRepo.CreateUser(user)
}
