package main

import (
	"fmt"
	"os"

	"github.com/AnnaRozhnova/blog"
	"github.com/AnnaRozhnova/blog/pkg/handler"
	"github.com/AnnaRozhnova/blog/pkg/repository"
	"github.com/AnnaRozhnova/blog/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)



func main() {

	// config
	if err := initConfig(); err != nil {
		fmt.Println("Error while initializing configs: ", err)
	}

	// loading .env file
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error while loading env file: ", err)
	}

	// connect to Postgres
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})


	// clean architecture: handler -> service -> repository

	// repository init
	repos := repository.NewRepository(db)
	// service init
	services := service.NewService(repos)
	// handler init
	handlers := handler.NewHandler(services)

	// server init
	srv := new(blog.Server)

	// run server
	err = srv.Run(os.Getenv("PORT"), handlers.InitRoutes())

	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
