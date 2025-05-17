package config

import (
	"errors"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/viper"
)

const (
	Development = "dev"
	Local       = "local"
	Production  = "prod"
)

var (
	errPathNotFound     = errors.New("path to config file not found")
	errFilenameNotFound = errors.New("name of config file not found")
)

type server struct {
	RTimeout time.Duration `mapstructure:"read_timeout"`
	WTimeout time.Duration `mapstructure:"write_timeout"`
	ITimeout time.Duration `mapstructure:"idle_timeout"`
	Host     string        `mapstructure:"host"`
	Port     string        `mapstructure:"port"`
}

type Config struct {
	Env    string `mapstructure:"env"`
	Server server `mapstructure:"server"`
}

func Load() (*Config, error) {
	path := os.Getenv("APP_CONFIG_PATH")
	if path == "" {
		return nil, errPathNotFound
	}
	filename := os.Getenv("APP_CONFIG_FILE")
	if filename == "" {
		return nil, errFilenameNotFound
	}
	fullpath := filepath.Join(path, filename)
	if _, err := os.Stat(fullpath); err != nil {
		return nil, err
	}

	v := viper.New()
	v.SetConfigFile(fullpath)

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := v.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
