package main

import (
	"log"

	"github.com/24zeroes/dmon"
	"github.com/24zeroes/dmon/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(dmon.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
