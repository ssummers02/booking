// Package bootstrap stores basic common entities such as config and database connection
package bootstrap

import (
	"errors"
	"fmt"
	"os"
	"reflect"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

const defaultEnvPath = ".env"

var ErrValidate = errors.New("bootstrap: validation failed")

type DBConfig struct {
	Scheme   string `envconfig:"DB_SCHEME"`
	Host     string `envconfig:"DB_HOST"`
	Port     string `envconfig:"DB_PORT"`
	Name     string `envconfig:"DB_NAME"`
	Username string `envconfig:"DB_USERNAME"`
	Password string `envconfig:"DB_PASSWORD"`
}

// Config is a common structure for all config types.
type Config struct {
	HTTPPort string `envconfig:"HTTP_PORT"`
	DB       DBConfig
}

// NewConfig loads configuration from the environment variables, optionally loading them from the file.
func NewConfig() (*Config, error) {
	err := godotenv.Load(defaultEnvPath)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return nil, err
	}

	var cfg Config

	err = envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}

	err = cfg.Validate()
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (c Config) Validate() error {
	s := reflect.ValueOf(c)
	if s.Kind() == reflect.Ptr {
		s = s.Elem()
	}

	var err string

	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		if f.IsZero() {
			err += " " + s.Type().Field(i).Name
		}
	}

	if err != "" {
		return fmt.Errorf("%w %s are empty", ErrValidate, err)
	}

	return nil
}
