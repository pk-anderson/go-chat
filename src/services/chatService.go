package services

import (
	"fmt"
	"log"
	"sort"

	"github.com/pk-anderson/go-chat/models"
	"github.com/pk-anderson/go-chat/repositories"
	"github.com/rabbitmq/amqp091-go"
)

type ChatService struct {
	rabbitMQURL string
	chatRepo    *repositories.ChatRepository
}

func generateQueueName(userID1, userID2 string) string {
	ids := []string{userID1, userID2}
	sort.Strings(ids)
	return fmt.Sprintf("chat_%s_%s", ids[0], ids[1])
}

func (cs *ChatService) SendMessage(senderID, receiverID, message string) error {
	queueName := generateQueueName(senderID, receiverID)

	existingChat, err := cs.chatRepo.FindChatByQueueName(queueName)
	if err != nil {
		return err
	}
	if existingChat == nil {
		conv := models.Chat{
			SenderID:   senderID,
			ReceiverID: receiverID,
			QueueName:  queueName,
		}
		err := cs.chatRepo.SaveChat(conv)
		if err != nil {
			return err
		}
	}

	conn, err := amqp091.Dial(cs.rabbitMQURL)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
		return err
	}
	defer ch.Close()

	_, err = ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
		return err
	}

	err = ch.Publish(
		"",
		queueName,
		false,
		false,
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	if err != nil {
		log.Fatalf("Failed to publish a message: %v", err)
		return err
	}

	log.Printf("Sent message: %s", message)
	return nil
}

func NewChatService(rabbitMQURL string, chatRepo *repositories.ChatRepository) *ChatService {
	return &ChatService{
		rabbitMQURL: rabbitMQURL,
		chatRepo:    chatRepo,
	}
}
