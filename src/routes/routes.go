package routes

import (
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func StartRoutes(router *mux.Router, mongoClient *mongo.Client, rabbitMQURL string) {
	StartUserRoutes(router, mongoClient)
	StartChatRoutes(router, mongoClient, rabbitMQURL)
}
