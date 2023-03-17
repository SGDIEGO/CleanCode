package repository

import (
	"database/sql"
	"fmt"

	"github.com/SGDIEGO/CleanCode/internal/domain/entity"
	"github.com/SGDIEGO/CleanCode/internal/port"
)

type AssignmentRepo struct {
	db *sql.DB
}

func NewAssignmentRepo(db *sql.DB) port.AssignmentRepo {
	return &AssignmentRepo{
		db: db,
	}
}

func (ar *AssignmentRepo) UserToGroup(assignmentEntity *entity.Assignment) error {

	query := `INSERT INTO assignmentT(userId, groupId, administrator) VALUES (?, ?, ?);`

	_, err := ar.db.Exec(query, assignmentEntity.UserId, assignmentEntity.GroupId, false)

	return err
}

func (ar *AssignmentRepo) DeleteAssignment(userId, groupId int) error {
	query := "DELETE FROM assignmentT WHERE groupId = ? AND userId = ?"

	response, err := ar.db.Exec(query, groupId, userId)

	if err != nil {
		return err
	}

	numberColumns, _ := response.RowsAffected()
	if numberColumns == 0 {
		return fmt.Errorf("no one in this group have number id %v", userId)
	}

	return nil
}
