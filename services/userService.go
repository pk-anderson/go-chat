package services

import (
	"github.com/pk-anderson/go-chat/models"
	"github.com/pk-anderson/go-chat/repositories"
	"github.com/pk-anderson/go-chat/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	repo *repositories.UserRepository
}

func (s *UserService) CreateUser(username, password string) (*models.User, error) {
	hashedPassword, err := utils.HashedPassword(password)
	if err != nil {
		return nil, err
	}

	user := models.User{
		Username: username,
		Password: string(hashedPassword),
	}

	result, err := s.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	user.ID = result.InsertedID.(primitive.ObjectID)

	return &user, nil
}

func (s *UserService) ListUsers() ([]models.User, error) {
	users, err := s.repo.ListUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}
