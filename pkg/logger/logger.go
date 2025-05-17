package logger

import (
	"errors"

	"github.com/DevAthhh/auth-service/pkg/config"
	"go.uber.org/zap"
)

func Load(env string) (*zap.Logger, error) {
	switch env {
	case config.Development, config.Local:
		return zap.NewDevelopment()
	case config.Production:
		return zap.NewProduction()
	}
	return nil, errors.New("unknown environment")
}
