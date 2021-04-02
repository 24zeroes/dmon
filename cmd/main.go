package main

import (
	"log"

	"github.com/24zeroes/dmon"
	"github.com/24zeroes/dmon/pkg/handler"
	"github.com/24zeroes/dmon/pkg/repository"
	"github.com/24zeroes/dmon/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(dmon.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
