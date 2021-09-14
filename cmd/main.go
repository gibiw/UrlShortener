package main

import (
	"log"

	server "github.com/gibiw/UrlShortener/internal"
	"github.com/gibiw/UrlShortener/internal/handler"
	"github.com/gibiw/UrlShortener/internal/repository"
	"github.com/gibiw/UrlShortener/internal/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(server.Server)

	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatal(err.Error())
	}
}
