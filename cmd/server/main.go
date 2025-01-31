package main

import (
	"github.com/Le0nar/calculate_xlsx/internal/handler"
	"github.com/Le0nar/calculate_xlsx/internal/service"
)

func main() {
	service := service.NewService()
	handler := handler.NewHandler(service)

	router := handler.InitRouter()

	router.Run()
}
