package services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/pk-anderson/go-chat/config"
	"github.com/pk-anderson/go-chat/models"
	"github.com/pk-anderson/go-chat/repositories"
	"github.com/pk-anderson/go-chat/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	repo     *repositories.UserRepository
	authRepo *repositories.AuthTokenRepository
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

func (s *UserService) Authenticate(id, password string) (string, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return "", err
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	token := jwt.New(jwt.SigningMethodHS256)
	expiresAt := time.Now().Add(time.Hour * 72).Unix()
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["exp"] = expiresAt

	t, err := token.SignedString([]byte(config.JWTSecretKey))
	if err != nil {
		return "", err
	}

	s.authRepo.CreateToken(models.AuthToken{
		Token:     t,
		UserID:    id,
		ExpiresAt: expiresAt,
	})

	return t, nil
}

func (s *UserService) ListUsers() ([]models.User, error) {
	users, err := s.repo.ListUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func NewUserService(repo *repositories.UserRepository, authRepo *repositories.AuthTokenRepository) *UserService {
	return &UserService{
		repo:     repo,
		authRepo: authRepo,
	}
}
