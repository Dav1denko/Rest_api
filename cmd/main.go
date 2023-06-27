package main

import (
	"log"
	restapi "rest_api_TODO"
	"rest_api_TODO/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(restapi.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while renning http server: %s", err.Error())
	}
}
