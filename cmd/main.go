package main

import (
	"net/http"

	"github.com/joel-kores/Conduit_Real-World/internal/config"
	"github.com/joel-kores/Conduit_Real-World/internal/database"
	"github.com/joel-kores/Conduit_Real-World/internal/generated"
	"github.com/joel-kores/Conduit_Real-World/internal/handlers"
	"github.com/joel-kores/Conduit_Real-World/internal/repository"
	"github.com/joel-kores/Conduit_Real-World/internal/services"
	"github.com/joel-kores/Conduit_Real-World/pkg/jwt"
	"github.com/joel-kores/Conduit_Real-World/pkg/logger"
)

func main() {
	cfg := config.LoadConfig()
	appLogger := logger.NewZeroLogger(cfg.LogLevel)

	db, err := database.NewDatabase(database.DatabaseConfig{
		Host: cfg.Database.Host,
		Port: cfg.Database.Port,
		User: cfg.Database.User,
		Password: cfg.Database.Password,
		DBName: cfg.Database.DBName,
		SSLMode: cfg.Database.SSLMode,
	})
	if err != nil {
		appLogger.Error("failed to connect to database", map[string]interface{}{"error": err.Error()})
		return
	}

	if err := database.AutoMigrate(db); err != nil {
		appLogger.Error("failed to run database migrations", map[string]interface{}{"error": err.Error()})
		return
	}

	userRepo := repository.NewPostgresUserRepository(db)

	jwtManager := jwt.NewJWTManager(cfg.JWTSecret, cfg.AccessTokenDur)

	authService := services.NewAuthService(services.AuthServiceConfig{
		UserRepo: userRepo,
		JWTManager: jwtManager,
		Logger: appLogger,
	})

	// authHandler := handlers.NewAuthHandler(authService, appLogger)

	handler := handlers.NewCompleteHandler(authService, appLogger)

	router := generated.Handler(handler)

	appLogger.Info("Starting server", map[string]interface{}{"port": cfg.ServerPort})
    http.ListenAndServe(cfg.ServerPort, router)
}

