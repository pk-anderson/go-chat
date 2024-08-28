package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pk-anderson/go-chat/config"
	"github.com/pk-anderson/go-chat/routes"
)

func main() {
	config.ConnectMongoDB("mongodb://localhost:27017")

	router := mux.NewRouter()

	routes.StartUserRoutes(router, config.MongoClient)

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
