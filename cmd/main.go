package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/pwkm/client/internal/core/util"
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
	db := setupDatabase(config)

	// userRepo := postgres.NewPostgresUserRepository(db) // Create the adapter
	db, err := postgres.New(ctx, config.DB)
	// userService := application.NewUserService(userRepo) // Inject into the service

	// userHandler := http.NewUserHandler(userService)     // Wire the handler

	// // Set up the HTTP server
	// http.HandleFunc("/users", userHandler.CreateUser)
	// log.Fatal(http.ListenAndServe(":8080", nil))

}

func setupDatabase(config *util.Container) *sql.DB {
	// -- Database config
	// Database connection
	dataSourceName := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.DB.Host, config.DB.Port, config.DB.User, config.DB.Password, config.DB.Name)

	db, err := sql.Open(config.DB.Connection, dataSourceName)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	return db

}
