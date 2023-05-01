package main

import (
	"log"

	"github.com/Marif226/go-todo-rest/internal/handler"
)

func main() {
	h := handler.New()

	// srv := server.New()
	e := h.InitRoutes()
	
	// err := srv.Run("8080", h.InitRoutes())
	err := e.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
}