package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/pk-anderson/go-chat/middlewares"
	"github.com/pk-anderson/go-chat/services"
)

type ChatHandler struct {
	chatService *services.ChatService
}

func (ch *ChatHandler) SendMessage(w http.ResponseWriter, r *http.Request) {
	var req struct {
		ReceiverID string `json:"receiverID"`
		Message    string `json:"message"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	senderID, ok := middlewares.GetUserIDFromContext(r.Context())
	if !ok {
		http.Error(w, "Unable to retrieve user ID", http.StatusUnauthorized)
		return
	}

	err := ch.chatService.SendMessage(senderID, req.ReceiverID, req.Message)
	if err != nil {
		http.Error(w, "Failed to send message", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Message sent successfully"))
}

func NewChatHandler(chatService *services.ChatService) *ChatHandler {
	return &ChatHandler{chatService: chatService}
}
