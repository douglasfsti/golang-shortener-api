package main

import (
	"github.com/douglasfsti/golang-shortener-api/config"
	"github.com/douglasfsti/golang-shortener-api/internal/api"
	"log"
	"net/http"
)

func main() {
	container := config.NewContainer()
	router := api.GetRouter(container)

	log.Fatal(http.ListenAndServe(":8080", router))
}
