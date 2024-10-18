package util

import (
	"github.com/spf13/viper"
)

// Container contains environment variables for the application, database, cache, token, and http server
type (
	Container struct {
		App   *App
		Token *Token
		Redis *Redis
		DB    *DB
		HTTP  *HTTP
	}
	// App contains all the environment variables for the application
	App struct {
		Name string
		Env  string
	}
	// Token contains all the environment variables for the token service
	Token struct {
		Duration string
	}
	// Redis contains all the environment variables for the cache service
	Redis struct {
		Addr     string
		Password string
	}
	// Database contains all the environment variables for the database
	DB struct {
		Connection string
		Host       string
		Port       string
		User       string
		Password   string
		Name       string
	}
	// HTTP contains all the environment variables for the http server
	HTTP struct {
		Env            string
		URL            string
		Port           string
		AllowedOrigins string
	}
)

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config *Container, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	app := &App{
		Name: viper.GetString("APP_NAME"),
		Env:  viper.GetString("APP_ENV"),
	}

	token := &Token{
		Duration: viper.GetString("TOKEN_DURATION"),
	}

	redis := &Redis{
		Addr:     viper.GetString("REDIS_ADDR"),
		Password: viper.GetString("REDIS_PASSWORD"),
	}

	db := &DB{
		Connection: viper.GetString("DB_CONNECTION"),
		Host:       viper.GetString("DB_HOST"),
		Port:       viper.GetString("DB_PORT"),
		User:       viper.GetString("DB_USER"),
		Password:   viper.GetString("DB_PASSWORD"),
		Name:       viper.GetString("DB_NAME"),
	}

	http := &HTTP{
		Env:            viper.GetString("APP_ENV"),
		URL:            viper.GetString("HTTP_URL"),
		Port:           viper.GetString("HTTP_PORT"),
		AllowedOrigins: viper.GetString("HTTP_ALLOWED_ORIGINS"),
	}

	return &Container{
		app,
		token,
		redis,
		db,
		http,
	}, nil

}
