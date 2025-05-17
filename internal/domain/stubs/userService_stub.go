package stubs

import (
	"github.com/DevAthhh/auth-service/internal/domain/models"
)

type UserService struct {
}

var (
	userRes = models.NewUser("Atheros", "ppp@pp.kk", "1111")
)

func NewUserService() *UserService {
	return &UserService{}
}

func (us *UserService) CreateUser(user *models.User) (*models.User, error) {
	return userRes, nil
}

func (us *UserService) FindUserByEmail(user *models.User) (*models.User, error) {
	return userRes, nil
}

func (us *UserService) DeleteUserByEmail(user *models.User) (*models.User, error) {
	return userRes, nil
}

func (us *UserService) ChangePassword(user *models.User) (*models.User, error) {
	return userRes, nil
}

func (us *UserService) ComparePassword(userDB *models.User, userRequest *models.User) error {
	return nil
}
