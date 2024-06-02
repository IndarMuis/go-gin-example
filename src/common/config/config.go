package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config interface {
	Get(key string) string
}

type configImpl struct {
	dotEnvExist bool
}

func (config *configImpl) Get(key string) string {
	osEnv := os.Getenv(key)
	if osEnv != "" {
		return osEnv
	}

	if !config.dotEnvExist {
		return ""
	}

	dotEnv, ok := os.LookupEnv(key)
	if !ok {
		return ""
	}

	return dotEnv
}

func New(filenames ...string) Config {
	err := godotenv.Load(filenames...)
	dotEnvExist := true
	if err != nil {
		dotEnvExist = false
	}
	return &configImpl{dotEnvExist: dotEnvExist}
}
