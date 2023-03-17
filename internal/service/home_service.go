package service

import (
	"github.com/SGDIEGO/CleanCode/internal/domain/entity"
	"github.com/SGDIEGO/CleanCode/internal/port"
)

type HomeService interface {
	Index() (*[]entity.User, error)
}

type homeService struct {
	UserRepo port.UserRepo
}

func NewHomeService(UserRepo port.UserRepo) HomeService {
	return &homeService{
		UserRepo: UserRepo,
	}
}

func (h *homeService) Index() (*[]entity.User, error) {

	users, err := h.UserRepo.GetAllUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}
