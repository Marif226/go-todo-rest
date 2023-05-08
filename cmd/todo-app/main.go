package main

import (
	"log"
	"os"

	"github.com/Marif226/go-todo-rest/internal/handler"
	"github.com/Marif226/go-todo-rest/internal/repository"
	"github.com/Marif226/go-todo-rest/internal/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	err := initConfig()
	if err != nil {
		log.Fatalf("error during reading the config file, %s", err.Error())
	}

	err = godotenv.Load()
	if err != nil {
		log.Fatalf("error during loading .env variable, %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host: viper.GetString("db.host"),
		Port: viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName: viper.GetString("db.dbname"),
		SSLMode: viper.GetString("db.sslmode"),
	})

	if err != nil {
		log.Fatalf("error during connecting to the database, %s", err.Error())
	}

	r := repository.New(db)
	s := service.New(r)
	h := handler.New(s)

	// srv := server.New()
	e := h.InitRoutes()
	
	// err := srv.Run("8080", h.InitRoutes())
	err = e.Start(viper.GetString("port"))
	if err != nil {
		log.Fatalf("error during running http server, %s", err.Error())
	}
}

func initConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("config")

	return viper.ReadInConfig()
}