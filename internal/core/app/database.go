package app

import (
	"context"
	"database/sql"
	"fmt"
	"net"

	"cloud.google.com/go/cloudsqlconn"
	"github.com/SGDIEGO/CleanCode/internal/core/config"
	"github.com/go-sql-driver/mysql"
)

type DataBase struct {
	DbConfig *config.DbConfig
}

func NewDatabase(DbConfig *config.DbConfig) *DataBase {

	return &DataBase{
		DbConfig: DbConfig,
	}
}

// Google cloud
func (db *DataBase) Load() (*sql.DB, error) {
	var (
		dbUser                 = db.DbConfig.User
		dbPwd                  = db.DbConfig.Password
		dbName                 = db.DbConfig.Name
		instanceConnectionName = db.DbConfig.Instance
		credentials            = db.DbConfig.Credentials
	)

	// Create a new dialer with authentication options
	// cloudsqlconn.WithCredentialsFile("C:\\Users\\51916\\Downloads\\helpful-quanta-378421-03e5fcee4235.json")
	d, err := cloudsqlconn.NewDialer(context.Background(), cloudsqlconn.WithCredentialsFile(credentials))
	if err != nil {
		return nil, fmt.Errorf("cloudsqlconn.NewDialer: %v", err)
	}

	var opts []cloudsqlconn.DialOption
	mysql.RegisterDialContext("cloudsqlconn",
		func(ctx context.Context, addr string) (net.Conn, error) {
			return d.Dial(ctx, instanceConnectionName, opts...)
		})

	dbURI := fmt.Sprintf("%s:%s@cloudsqlconn(localhost:3306)/%s?parseTime=true",
		dbUser, dbPwd, dbName)

	dbPool, err := sql.Open("mysql", dbURI)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %v", err)
	}
	return dbPool, nil
}
