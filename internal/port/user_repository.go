package port

import "github.com/SGDIEGO/CleanCode/internal/domain/entity"

type UserRepo interface {
	GetAllUsers() (*[]entity.User, error)
	GetUserById(id int) (*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
	CreateUser(user *entity.User) error
	UpdateUser(id int, user *entity.User) error
	DeleteUser(id int) error
}
