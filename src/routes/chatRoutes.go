package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pk-anderson/go-chat/handlers"
	"github.com/pk-anderson/go-chat/middlewares"
	"github.com/pk-anderson/go-chat/repositories"
	"github.com/pk-anderson/go-chat/services"
	"go.mongodb.org/mongo-driver/mongo"
)

func StartChatRoutes(router *mux.Router, mongoClient *mongo.Client, rabbitMQURL string) {
	chatRepo := repositories.NewChatRepository(mongoClient, "chat_db", "chat")
	chatService := services.NewChatService(rabbitMQURL, chatRepo)
	chatHandler := handlers.NewChatHandler(chatService)

	protectedRouter := router.PathPrefix("/api").Subrouter()
	protectedRouter.Use(middlewares.JWTMiddleware)

	protectedRouter.HandleFunc("/chat/send", chatHandler.SendMessage).Methods(http.MethodPost)
}
