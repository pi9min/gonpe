package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/pi9min/gonpe/server/application"
	"github.com/pi9min/gonpe/server/application/repository"
	"github.com/pi9min/gonpe/server/config"
	"github.com/pi9min/gonpe/server/handler"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func init() {
	config.Setup()

	repos := repository.NewAllRepository()
	authApp := application.NewAdminApp(
		repos,
	)

	http.Handle("/", handler.NewHandler(authApp))
}
