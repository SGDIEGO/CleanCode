package app

import (
	"fmt"
	"log"

	"github.com/SGDIEGO/CleanCode/internal/core/config"
	"github.com/SGDIEGO/CleanCode/internal/domain/repository"
)

type App struct {
	Server   *Server
	DataBase *DataBase
}

func NewApp() *App {
	// Load configuration
	config, err := config.LoadConfig()

	if err != nil {
		log.Fatal(err.Error())
		return &App{
			// Config:   nil,
			Server:   nil,
			DataBase: nil,
		}
	}

	server := NewServer(config.SvConfig)
	database := NewDatabase(config.DbConfig)

	return &App{
		Server:   server,
		DataBase: database,
	}
}

func (a *App) RunApp() {

	// Load DB
	db, err := a.DataBase.Load()

	if err != nil {
		fmt.Println("Error loading database:", err.Error())
	}

	defer db.Close()

	// Initizalize tables if not exists
	if err := repository.InitTables(db); err != nil {
		fmt.Println("Error initialising tables", err.Error())
	}

	// Load Server
	a.Server.Load(db)
}
