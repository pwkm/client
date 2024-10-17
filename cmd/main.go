package main

import (
	"context"
	"database/sql"
	"log"

	"gorm.io/driver/postgres"
)

func main() {
	// Load environment variables
	config, err := config.New()
	if err != nil {
		log.Fatal("Error loading environment variables", err)
	}

	// Set logger
	// logger.Set(config.App)
	//slog.Info("Starting the application", "app", config.App.Name, "env", config.App.Env)

	// Setting the context
	ctx := context.Background()

	// -- Database config
	db, err := postgres.New(ctx, config.DB)
	db, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	// db := setupPostgresDB() // Connect to PostgreSQL

	// userRepo := postgres.NewPostgresUserRepository(db) // Create the adapter
	// userService := application.NewUserService(userRepo) // Inject into the service

	// userHandler := http.NewUserHandler(userService)     // Wire the handler

	// // Set up the HTTP server
	// http.HandleFunc("/users", userHandler.CreateUser)
	// log.Fatal(http.ListenAndServe(":8080", nil))

}
