package port

import "github.com/SGDIEGO/CleanCode/internal/domain/entity"

type GroupRepo interface {
	GetAllGroups() (*[]entity.Group, error)
	GetGroupById(id int) (*entity.Group, error)
	GetGroupsByAdminName(adminName string) (*[]entity.Group, error)
	CreateGroup(group *entity.Group) error
	UpdateGroup(id int, group *entity.Group) error
	DeleteGroup(id int) error
	UserIdsFromGroup(id int) ([]int, error)
}
