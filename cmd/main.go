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



	if err := initConfig(); err != nil {
		fmt.Println("Error while initializing configs: ", err)
	}
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error while loading env file: ", err)
	}


	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(blog.Server)
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
