package repositories

import (
	"context"
	"fmt"

	"github.com/pk-anderson/go-chat/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

func (r *UserRepository) CreateUser(user models.User) (*mongo.InsertOneResult, error) {
	return r.collection.InsertOne(context.TODO(), user)
}

func (r *UserRepository) ListUsers() ([]models.User, error) {
	var users []models.User

	cursor, err := r.collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return nil, fmt.Errorf("failed to find users: %v", err)
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			return nil, fmt.Errorf("failed to decode user: %v", err)
		}
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}

	return users, nil
}

func NewUserRepository(client *mongo.Client, dbName, collectionName string) *UserRepository {
	collection := client.Database(dbName).Collection(collectionName)
	return &UserRepository{collection: collection}
}
