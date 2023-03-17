package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/SGDIEGO/CleanCode/internal/domain/entity"
	"github.com/SGDIEGO/CleanCode/internal/port"
)

type UserRepo struct {
	db *sql.DB
}

func NewuserRepo(db *sql.DB) port.UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (u *UserRepo) GetAllUsers() (*[]entity.User, error) {

	getUsers := "SELECT * FROM user;"
	response, err := u.db.Query(getUsers)

	if err != nil {
		return nil, err
	}

	var user entity.User
	var users []entity.User

	for response.Next() {
		response.Scan(&user.UserId, &user.UserName, &user.Email, &user.Password)
		users = append(users, user)
	}

	return &users, nil
}

func (u *UserRepo) GetUserById(id int) (*entity.User, error) {
	getUser := "SELECT * FROM user WHERE userId = ?;"
	response := u.db.QueryRow(getUser, id)

	if response.Err() != nil {
		return nil, response.Err()
	}

	var user entity.User

	if err := response.Scan(&user.UserId, &user.UserName, &user.Email, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no se encontró ningún usuario con ID %d", id)
		}
		return nil, err
	}

	return &user, nil
}

func (u *UserRepo) GetUserByEmail(email string) (*entity.User, error) {
	getUser := "SELECT * FROM user WHERE email = ?;"
	response := u.db.QueryRow(getUser, email)

	var user entity.User

	if err := response.Scan(&user.UserId, &user.UserName, &user.Email, &user.Password); err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserRepo) CreateUser(user *entity.User) error {

	existsUser := existRow(u.db, "user", "email", user.Email)

	if existsUser {
		return errors.New("user with this email exists")
	}

	query := `INSERT INTO user(userName, email, password) VALUES( ?, ?, ?)`

	_, err := u.db.Exec(query, user.UserName, user.Email, user.Password)

	return err
}

func (u *UserRepo) UpdateUser(id int, user *entity.User) error {

	existsUser := existRow(u.db, "user", "userID", id)

	if !existsUser {
		return errors.New("user with this id doesn't exists")
	}

	query := `
			UPDATE user 
			SET 
				userName = ?,
				email = ?,
				password = ?,
			WHERE
				userId = ?;
			`

	_, err := u.db.Exec(query, user.UserName, user.Email, user.Password, id)

	return err
}

func (u *UserRepo) DeleteUser(id int) error {
	existsUser := existRow(u.db, "user", "userID", id)

	if !existsUser {
		return errors.New("user with this id doesn't exists")
	}

	query := `DELETE FROM user 	WHERE userId = ?`

	_, err := u.db.Exec(query, id)

	return err
}
