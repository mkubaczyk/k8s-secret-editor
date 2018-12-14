package settings

import (
	"github.com/apsdehal/go-logger"
	"os"
)

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

var Logger *logger.Logger

func Init() {
	logger.SetDefaultFormat("%{time:2006-01-02 15:04:05} [%{level}] %{message}")
	Logger, _ = logger.New("logger", 1, os.Stdout)
	Logger.Info("Config initialized")
}
