package main

import (
	"log"

	"github.com/Le0nar/calculate_xlsx/internal/handler"
	"github.com/Le0nar/calculate_xlsx/internal/repository"
	"github.com/Le0nar/calculate_xlsx/internal/service"
)

func main() {
	db, err := repository.NewDB()

	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	defer db.Close()

	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	handler := handler.NewHandler(service)

	router := handler.InitRouter()

	router.Run()
}
