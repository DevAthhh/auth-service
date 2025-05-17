package services

import (
	"context"

	"github.com/DevAthhh/auth-service/internal/domain/models"
	"github.com/DevAthhh/auth-service/internal/domain/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceInterface interface {
	CreateUser(user *models.User) (*models.User, error)
	FindUserByEmail(user *models.User) (*models.User, error)
	DeleteUserByEmail(user *models.User) (*models.User, error)
	ChangePassword(user *models.User) (*models.User, error)
	ComparePassword(userDB *models.User, userRequest *models.User) error
}

type UserService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (us *UserService) CreateUser(user *models.User) (*models.User, error) {
	return us.userRepo.CreateUser(context.Background(), user)
}

func (us *UserService) FindUserByEmail(user *models.User) (*models.User, error) {
	return us.userRepo.GetUserByEmail(context.Background(), user)
}

func (us *UserService) DeleteUserByEmail(user *models.User) (*models.User, error) {
	return us.userRepo.DeleteUserByEmail(context.Background(), user)
}

func (us *UserService) ChangePassword(user *models.User) (*models.User, error) {
	return us.userRepo.ChangePasswordByEmail(context.Background(), user)
}

func (us *UserService) ComparePassword(userDB *models.User, userRequest *models.User) error {
	return bcrypt.CompareHashAndPassword([]byte(userDB.GetPassword()), []byte(userRequest.GetPassword()))
}
