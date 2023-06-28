package main

import (
	"log"
	restapi "rest_api_TODO"
	"rest_api_TODO/pkg/handler"
	"rest_api_TODO/pkg/repository"
	"rest_api_TODO/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(restapi.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while renning http server: %s", err.Error())
	}
}
