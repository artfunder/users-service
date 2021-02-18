package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/artfunder/users-service/repository"
	"github.com/artfunder/users-service/router"
	"github.com/artfunder/users-service/service"
)

func main() {
	r := router.CreateRouter(service.NewUsersService(repository.NewSQLiteRepo()))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("APP_PORT")), r))
}
