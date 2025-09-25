package services

import (
	"github.com/joel-kores/Conduit_Real-World/internal/domain/user"
	"github.com/joel-kores/Conduit_Real-World/pkg/jwt"
	"github.com/joel-kores/Conduit_Real-World/pkg/logger"
)

type AuthService struct {
	userRepo   user.UserRepository
	jwtManager *jwt.JWTManager
	logger     logger.Logger
}

type AuthServiceConfig struct {
	UserRepo   user.UserRepository
	JWTManager *jwt.JWTManager
	Logger     logger.Logger
}

func NewAuthService(cfg AuthServiceConfig) *AuthService {
	return &AuthService{
		userRepo:   cfg.UserRepo,
		jwtManager: cfg.JWTManager,
		logger:     cfg.Logger,
	}
}

func (s AuthService) Register(req user.RegisterRequest) (*user.User, error) {
	s.logger.Info("Registering new user", map[string]interface{}{"email": req.Email})
	exists, err := s.userRepo.UserExistsByEmail(req.Email)
	if err != nil {
		s.logger.Error("Error checking if user exists", map[string]interface{}{"error": err.Error()})
		return nil, err
	}
	if exists {
		s.logger.Warn("User already exists", map[string]interface{}{"email": req.Email})
		return nil, user.ErrUserAlreadyExists
	}

	newUser := &user.User{
		Username: req.UserName,
		Email:    req.Email,
	}

	if err := newUser.SetPassword(req.Password); err != nil {
		s.logger.Error("Error setting password", map[string]interface{}{"error": err.Error()})
		return nil, err
	}

	if err := s.userRepo.CreateUser(newUser); err != nil {
		s.logger.Error("Error creating user", map[string]interface{}{"error": err.Error()})
		return nil, err
	}

	s.logger.Info("User registered successfully", map[string]interface{}{"user_id": newUser.ID})
	return newUser, nil
}

func (s *AuthService) Login(req user.LoginRequest) (*user.User, error) {
	s.logger.Info("user login attempt", map[string]interface{}{"email": req.Email})
	user, err := s.userRepo.FindUserByEmail(req.Email)
	if err != nil {
		s.logger.Warn("User not found", map[string]interface{}{"email": req.Email})
		return nil, err
	}
	if !user.CheckPassword(req.Password) {
		s.logger.Warn("invalid password", map[string]interface{}{"email": req.Email})
		return nil, err
	}
	s.logger.Info("user logged in successfully", map[string]interface{}{"user_id": user.ID})
	return user, nil
}
