package repository

import (
	"database/sql"
	"errors"

	"github.com/SGDIEGO/CleanCode/internal/domain/entity"
	"github.com/SGDIEGO/CleanCode/internal/port"
)

type GroupRepo struct {
	db *sql.DB
}

func NewGroupRepo(db *sql.DB) port.GroupRepo {
	return &GroupRepo{
		db: db,
	}
}

func (g *GroupRepo) GetAllGroups() (*[]entity.Group, error) {
	getGroups := "SELECT * FROM groupU;"
	response, err := g.db.Query(getGroups)

	if err != nil {
		return nil, err
	}

	var group entity.Group
	var groups []entity.Group

	for response.Next() {
		response.Scan(&group.GroupId, &group.GroupName, &group.Admin, &group.Members, &group.Created)
		groups = append(groups, group)
	}
	return &groups, nil
}

func (g *GroupRepo) GetGroupById(id int) (*entity.Group, error) {
	getGroups := "SELECT * FROM groupU WHERE groupId = ?"
	response := g.db.QueryRow(getGroups, id)

	var group entity.Group

	if err := response.Scan(&group.GroupId, &group.GroupName, &group.Admin, &group.Members, &group.Created); err != nil {
		return nil, err
	}

	return &group, nil
}

func (g *GroupRepo) GetGroupsByAdminName(adminName string) (*[]entity.Group, error) {

	queryGroups := "SELECT * FROM groupU WHERE admin = ?"

	response, err := g.db.Query(queryGroups, adminName)

	if err != nil {
		return nil, err
	}

	var Group entity.Group
	var Groups []entity.Group

	for response.Next() {
		if err := response.Scan(&Group.GroupId, &Group.GroupName, &Group.Admin, &Group.Members, &Group.Created); err != nil {
			return nil, err
		}
		Groups = append(Groups, Group)
	}

	return &Groups, nil
}

func (g *GroupRepo) CreateGroup(group *entity.Group) error {

	existsGroup := existRow(g.db, "groupU", "groupName", group.GroupName)

	if existsGroup {
		return errors.New("group with this email exists")
	}

	query := `INSERT INTO groupU(groupName, admin, members, created) VALUES( ?, ?, ?, ?)`

	_, err := g.db.Exec(query, group.GroupName, group.Admin, group.Members, group.Created)

	return err
}

func (g *GroupRepo) UpdateGroup(id int, group *entity.Group) error {
	existsGroup := existRow(g.db, "groupU", "groupName", id)

	if !existsGroup {
		return errors.New("group with this id doesn't exists")
	}

	query := `
			UPDATE groupU 
			SET 
				groupName = ?,
				admin = ?,
				members = ?,
				created = ?,
			WHERE
				groupId = ?;
			`

	_, err := g.db.Exec(query, group.GroupName, group.Admin, group.Members, group.Created, id)

	return err
}

func (g *GroupRepo) DeleteGroup(id int) error {
	existsGroup := existRow(g.db, "groupU", "groupId", id)

	if !existsGroup {
		return errors.New("group with this id doesn't exists")
	}

	query := `DELETE FROM groupU WHERE groupId = ?`

	_, err := g.db.Exec(query, id)

	return err
}

func (g *GroupRepo) UserIdsFromGroup(id int) ([]int, error) {
	query := "SELECT (userId) FROM assignmentT WHERE groupId = ?"

	result, err := g.db.Query(query, id)
	if err != nil {
		return nil, err
	}

	var UsersId []int
	var UserId int

	for result.Next() {
		if err := result.Scan(&UserId); err != nil {
			return nil, err
		}
		UsersId = append(UsersId, UserId)
	}

	return UsersId, nil
}
