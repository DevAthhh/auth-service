package persistance

import (
	"context"
	"time"

	"github.com/DevAthhh/auth-service/internal/domain/models"
	"github.com/DevAthhh/auth-service/internal/infrastructure/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func (ur *UserRepository) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 4*time.Second)
	defer cancel()

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.GetPassword()), 10)
	if err != nil {
		return nil, err
	}

	userDB := entity.User{
		Username: user.GetUsername(),
		Email:    user.GetEmail(),
		UUID:     user.GetID(),
		Password: string(passwordHash),
	}

	if err := ur.db.WithContext(ctx).Create(&userDB).Error; err != nil {
		return nil, err
	}

	resUser := models.NewUser(userDB.Username, userDB.Email, userDB.Password)

	return resUser, nil
}

func (ur *UserRepository) GetUserByEmail(ctx context.Context, user *models.User) (*models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 4*time.Second)
	defer cancel()

	var userDB entity.User
	if err := ur.db.WithContext(ctx).First(&userDB, "email = ?", user.GetEmail()).Error; err != nil {
		return nil, err
	}

	resUser := models.NewUser(userDB.Username, userDB.Email, userDB.Password)

	return resUser, nil

}

func (ur *UserRepository) DeleteUserByEmail(ctx context.Context, user *models.User) (*models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 4*time.Second)
	defer cancel()

	var userDB entity.User
	if err := ur.db.WithContext(ctx).Delete(&userDB, "email = ?", user.GetEmail()).Error; err != nil {
		return nil, err
	}

	resUser := models.NewUser(userDB.Username, userDB.Email, userDB.Password)

	return resUser, nil
}

func (ur *UserRepository) ChangePasswordByEmail(ctx context.Context, user *models.User) (*models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 4*time.Second)
	defer cancel()

	var userDB entity.User
	if err := ur.db.WithContext(ctx).First(&userDB, "email = ?", user.GetEmail()).Error; err != nil {
		return nil, err
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.GetPassword()), 10)
	if err != nil {
		return nil, err
	}

	userDB.Password = string(passwordHash)
	if err := ur.db.WithContext(ctx).Save(&user).Error; err != nil {
		return nil, err
	}

	resUser := models.NewUser(userDB.Username, userDB.Email, userDB.Password)

	return resUser, nil
}

func (ur *UserRepository) IsEmailExists(ctx context.Context, user *models.User) error {
	if err := ur.db.First(&entity.User{}, "email = ?", user.GetEmail()).Error; err != nil {
		return err
	}
	return nil
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}
