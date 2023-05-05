package main

import (
	"log"

	"github.com/Marif226/go-todo-rest/internal/handler"
	"github.com/Marif226/go-todo-rest/internal/repository"
	"github.com/Marif226/go-todo-rest/internal/service"
	"github.com/spf13/viper"
)

func main() {
	err := initConfig()
	if err != nil {
		log.Fatalf("error during reading the config file, %s", err.Error())
	}

	r := repository.New()
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