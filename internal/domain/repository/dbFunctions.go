package repository

import (
	"database/sql"
)

func existRow(dB *sql.DB, table string, column string, option interface{}) bool {

	user := "SELECT COUNT(*) FROM ? WHERE ? = ?"

	row := dB.QueryRow(user, table, column, option)

	var numbeRows int

	row.Scan(&numbeRows)

	return numbeRows > 0
}

// Execute the queryString and catch the error
func createTable(queryString string, db *sql.DB) error {
	_, err := db.Exec(queryString)

	return err
}

func InitTables(db *sql.DB) error {

	UserT := `
		CREATE TABLE IF NOT EXISTS user(
			userId INT PRIMARY KEY AUTO_INCREMENT,
			userName VARCHAR(20) NOT NULL,
			email    VARCHAR(20) NOT NULL,
			password VARCHAR(20) NOT NULL
		);
	`

	GroupT := `
		CREATE TABLE IF NOT EXISTS groupU(
			groupId INT PRIMARY KEY AUTO_INCREMENT,
			groupName VARCHAR(20) NOT NULL,
			admin    VARCHAR(20) NOT NULL,
			members INT NOT NULL,
			created DATE NOT NULL
		);
	`

	AssignmentT := `
	CREATE TABLE IF NOT EXISTS assignmentT(
		assignId INT PRIMARY KEY AUTO_INCREMENT,
		userId INT,
		groupId INT,
		administrator BOOLEAN,
		FOREIGN KEY (userId) REFERENCES user(userId),
		FOREIGN KEY (groupId) REFERENCES groupU(groupId)
	);
`
	if err := createTable(UserT, db); err != nil {
		return err
	}

	if err := createTable(GroupT, db); err != nil {
		return err
	}

	if err := createTable(AssignmentT, db); err != nil {
		return err
	}

	return nil
}
