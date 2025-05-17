package app

import (
	"log"
	"os"

	"github.com/DevAthhh/auth-service/internal/domain/services"
	"github.com/DevAthhh/auth-service/internal/domain/stubs"
	"github.com/DevAthhh/auth-service/internal/infrastructure/database"
	server "github.com/DevAthhh/auth-service/internal/infrastructure/http"
	"github.com/DevAthhh/auth-service/internal/infrastructure/http/handlers"
	"github.com/DevAthhh/auth-service/internal/infrastructure/persistance"
	"github.com/DevAthhh/auth-service/pkg/config"
	"github.com/DevAthhh/auth-service/pkg/loadenv"
	zlog "github.com/DevAthhh/auth-service/pkg/logger"
)

func Run() {
	if err := loadenv.Load(); err != nil {
		log.Fatal(err)
	}
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	logger, err := zlog.Load(cfg.Env)
	if err != nil {
		log.Fatal(err)
	}

	var userService services.UserServiceInterface

	switch cfg.Env {
	case config.Development:
		userService = stubs.NewUserService()
	case config.Local, config.Production:
		db, err := database.LoadDB()
		if err != nil {
			log.Fatal(err)
		}
		if err := database.SyncDB(db); err != nil {
			log.Fatal(err)
		}

		repo := persistance.NewUserRepository(db)
		userService = services.NewUserService(repo)
	}

	authService := services.NewAuthService(os.Getenv("SECRET_KEY"))

	routes := handlers.NewController(userService, *authService)
	h, err := server.NewHandler(cfg, routes)
	if err != nil {
		log.Fatal(err)
	}
	server := server.NewServer(cfg, h)
	logger.Info("server has been started")
	server.Start()
}
