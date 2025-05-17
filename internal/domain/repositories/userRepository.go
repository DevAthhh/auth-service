package repositories

import (
	"context"

	"github.com/DevAthhh/auth-service/internal/domain/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	GetUserByEmail(ctx context.Context, user *models.User) (*models.User, error)
	DeleteUserByEmail(ctx context.Context, user *models.User) (*models.User, error)
	ChangePasswordByEmail(ctx context.Context, user *models.User) (*models.User, error)
	IsEmailExists(ctx context.Context, user *models.User) error
}
