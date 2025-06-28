package main

import (
	"log"

	sandbox "github.com/YaMyha/glowing-adventure"
	"github.com/YaMyha/glowing-adventure/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(sandbox.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occurred while running http server: %s", err.Error())
	}
}
