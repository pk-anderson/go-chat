package repositories

import (
	"context"

	"github.com/pk-anderson/go-chat/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthTokenRepository struct {
	collection *mongo.Collection
}

func (r *AuthTokenRepository) CreateToken(token models.AuthToken) (*mongo.InsertOneResult, error) {
	return r.collection.InsertOne(context.TODO(), token)
}

func (r *AuthTokenRepository) FindToken(token string) (*models.AuthToken, error) {
	var result models.AuthToken
	err := r.collection.FindOne(context.TODO(), map[string]string{"token": token}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func NewAuthTokenRepository(client *mongo.Client, dbName, collectionName string) *AuthTokenRepository {
	collection := client.Database(dbName).Collection(collectionName)
	return &AuthTokenRepository{collection: collection}
}
