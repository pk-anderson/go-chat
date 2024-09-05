package repositories

import (
	"context"

	"github.com/pk-anderson/go-chat/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ChatRepository struct {
	collection *mongo.Collection
}

func NewChatRepository(client *mongo.Client, dbName, collectionName string) *ChatRepository {
	return &ChatRepository{
		collection: client.Database(dbName).Collection(collectionName),
	}
}

func (repo *ChatRepository) SaveChat(conv models.Chat) error {
	_, err := repo.collection.InsertOne(context.Background(), conv)
	return err
}

func (repo *ChatRepository) FindChatByQueueName(queueName string) (*models.Chat, error) {
	var conv models.Chat
	filter := bson.M{"queue_name": queueName}
	err := repo.collection.FindOne(context.Background(), filter).Decode(&conv)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &conv, nil
}
