package config

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	ServerPort      string
	JWTSecret       string
	AccessTokenDur  time.Duration
	RefreshTokenDur time.Duration
	LogLevel        string
	Database        DatabaseConfig
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func LoadConfig() *Config {
	accessTokenDur, _ := strconv.Atoi(getEnv("CONDUIT_ACCESS_TOKEN_DURATION", "15"))
	refreshTokenDuration, _ := strconv.Atoi(getEnv("CONDUIT_REFRESH_TOKEN_DURATION", "1440"))

	return &Config{
		ServerPort:      getEnv("CONDUIT_SERVER_PORT", ":8080"),
		JWTSecret:       getEnv("CONDUIT_JWT_SECRET", "your-secret-key=must-be-really-secure"),
		AccessTokenDur:  time.Duration(accessTokenDur) * time.Minute,
		RefreshTokenDur: time.Duration(refreshTokenDuration) * time.Minute,
		LogLevel:        getEnv("CONDUIT_LOG_LEVEL", "info"),
		Database: DatabaseConfig{
			Host:     getEnv("CONDUIT_DB_HOST", "localhost"),
			Port:     getEnv("CONDUIT_DB_PORT", "5432"),
			User:     getEnv("CONDUIT_DB_USER", "conduituser"),
			Password: getEnv("CONDUIT_DB_PASSWORD", "conduitpass"),
			DBName:   getEnv("CONDUIT_DB_NAME", "conduitdb"),
			SSLMode:  getEnv("CONDUIT_DB_SSLMODE", "disable"),
		},
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
