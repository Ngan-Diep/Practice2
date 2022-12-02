package main

import (
	"context"
	"fmt"
	"github.com/rs/cors"
	"log"
	"nails-backend/pkg/common"
	"nails-backend/pkg/handlers"
	"nails-backend/pkg/logging"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func ServerStart(router *mux.Router) {
	fmt.Println("Server started at http://0.0.0.0:8080")
	c := cors.New(cors.Options{
		AllowedMethods: []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		AllowedHeaders: []string{"token", "shop-id", "content-type"},
	})
	err := http.ListenAndServe(":8080", c.Handler(router))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	ServiceName := "nails_backend"

	logger := logging.New(logging.DEBUG, os.Stdout).WithField(logging.TagKey, ServiceName)
	ctx := logging.WithLogger(context.Background(), logger)

	common.SecretKey = os.Getenv("JWT_SECRET")
	if common.SecretKey == "" {
		logger.Fatalln("JWT_SECRET required")
	}

	router := handlers.CreateRouter(ctx)
	ServerStart(router)
}
