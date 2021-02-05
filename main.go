package main

import (
	"log"
	"net/http"

	"github.com/artfunder/users-service/repository"
	"github.com/artfunder/users-service/router"
	"github.com/artfunder/users-service/service"
)

func main() {
	r := router.CreateRouter(service.NewUsersService(repository.NewSQLiteRepo()))

	log.Fatal(http.ListenAndServe(":8080", r))
}
