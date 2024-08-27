package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pk-anderson/go-chat/handlers"
	"github.com/pk-anderson/go-chat/repositories"
	"github.com/pk-anderson/go-chat/services"
	"go.mongodb.org/mongo-driver/mongo"
)

func StartUserRoutes(router *mux.Router, mongoClient *mongo.Client) {
	userRepo := repositories.NewUserRepository(mongoClient, "chat_db", "users")
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	router.HandleFunc("/user/create", userHandler.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/user/list", userHandler.ListUsers).Methods(http.MethodGet)
}
