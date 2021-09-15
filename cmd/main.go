package main

import (
	"log"

	server "github.com/gibiw/UrlShortener/internal"
	"github.com/gibiw/UrlShortener/internal/handler"
	"github.com/gibiw/UrlShortener/internal/repository"
	"github.com/gibiw/UrlShortener/internal/service"
	"github.com/spf13/viper"
)

func main() {

	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	port := viper.GetString("port")

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services, viper.GetString("host"), port)
	srv := new(server.Server)

	if err := srv.Run(port, handlers.InitRoutes()); err != nil {
		log.Fatal(err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
