package server

import (
	"errors"
	"time"

	"github.com/DevAthhh/auth-service/internal/infrastructure/http/handlers"
	"github.com/DevAthhh/auth-service/pkg/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewHandler(cfg *config.Config, userController *handlers.Controller) (*gin.Engine, error) {
	var router *gin.Engine
	switch cfg.Env {
	case config.Development, config.Local:
		gin.SetMode(gin.DebugMode)
		router = gin.Default()
	case config.Production:
		gin.SetMode(gin.ReleaseMode)
		router = gin.New()
		router.Use(gin.Recovery())

		router.Use(cors.New(cors.Config{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{"GET", "POST"},
			AllowHeaders:     []string{"Authorization", "Content-Type"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}))
	}

	if router == nil {
		return nil, errors.New("unknown environment")
	}

	router.POST("/register", userController.NewRegister())
	router.POST("/login", userController.NewLogin())
	router.POST("/refresh", userController.RefreshToken())
	router.POST("/check", userController.ValidateToken())

	return router, nil
}
