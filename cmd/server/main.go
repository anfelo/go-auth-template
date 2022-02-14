package main

import (
	"net/http"

	"github.com/anfelo/go-auth-template/internal/database"
	transportHTTP "github.com/anfelo/go-auth-template/internal/transport/http"
	"github.com/anfelo/go-auth-template/internal/users"

	log "github.com/sirupsen/logrus"
)

// App - contain application information
type App struct {
	Name    string
	Version string
}

// Run - sets up our application
func (a *App) Run() error {
	log.SetFormatter(&log.JSONFormatter{})
	log.WithFields(
		log.Fields{
			"AppName":    a.Name,
			"AppVersion": a.Version,
		}).Info("Setting up application")

	var err error
	db, err := database.NewDatabase()
	if err != nil {
		return err
	}
	err = database.MigrateDB(db)
	if err != nil {
		return err
	}

	usersService := users.NewService(db)

	handler := transportHTTP.NewHandler(usersService)
	handler.SetupRoutes()

	if err := http.ListenAndServe(":3000", handler.Router); err != nil {
		log.Error("Failed to set up server")
		return err
	}

	return nil
}

func main() {
	log.Info("Go Auth Service")
	app := App{
		Name:    "Auth Service",
		Version: "1.0.0",
	}

	if err := app.Run(); err != nil {
		log.Error("Error starting up the auth service")
		log.Fatal(err)
	}
}
