package main

import (
	"context"
	"fmt"
	"log"

	"github.com/pwkm/client/internal/adapter/postgres"
	"github.com/pwkm/client/internal/util"
	"gorm.io/driver/postgres"
)

func main() {
	// Load environment variables
	config, err := util.LoadConfig("./config")
	if err != nil {
		log.Fatal("Error loading environment variables", err)
	}

	// Set logger
	// logger.Set(config.App)
	//slog.Info("Starting the application", "app", config.App.Name, "env", config.App.Env)
	fmt.Printf("Starting the application", "app", config.App.Name, "env", config.App.Env)

	// Setting the context
	ctx := context.Background()

	// -- Database config
	db, err := postgres.PostgresNew(config)
	defer db.Close()

	// userRepo := postgres.NewPostgresUserRepository(db) // Create the adapter

	// userService := application.NewUserService(userRepo) // Inject into the service

	// userHandler := http.NewUserHandler(userService)     // Wire the handler

	// // Set up the HTTP server
	// http.HandleFunc("/users", userHandler.CreateUser)
	// log.Fatal(http.ListenAndServe(":8080", nil))

}
