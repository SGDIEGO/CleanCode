package port

import "github.com/SGDIEGO/CleanCode/internal/domain/entity"

type AssignmentRepo interface {
	UserToGroup(assignmentEntity *entity.Assignment) error
	DeleteAssignment(userId, group int) error
}
