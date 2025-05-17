package handlers

import (
	"github.com/DevAthhh/auth-service/internal/domain/services"
)

type Controller struct {
	userService services.UserServiceInterface
	authService services.AuthService
}

func NewController(userService services.UserServiceInterface, authService services.AuthService) *Controller {
	return &Controller{
		userService: userService,
		authService: authService,
	}
}
