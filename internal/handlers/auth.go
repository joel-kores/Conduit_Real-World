package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/joel-kores/Conduit_Real-World/internal/converter"
	"github.com/joel-kores/Conduit_Real-World/internal/domain/user"
	"github.com/joel-kores/Conduit_Real-World/internal/generated"
	"github.com/joel-kores/Conduit_Real-World/internal/services"
	"github.com/joel-kores/Conduit_Real-World/pkg/logger"
)

type AuthHandler struct {
	AuthService *services.AuthService
	Logger      logger.Logger
}

func NewAuthHandler(authService *services.AuthService, logger logger.Logger) *AuthHandler {
	return &AuthHandler{
		AuthService: authService,
		Logger:      logger,
	}
}

func (h *AuthHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req generated.NewUser
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	mapReq := user.RegisterRequest{
		Email: req.Email,
		UserName: req.Username,
		Password: req.Password,
	}

	u, err := h.AuthService.Register(mapReq)

	if err != nil {
		switch err {
		case user.ErrUserAlreadyExists:
			errorResponse := generated.GenericErrorModel{}
			errorResponse.Errors.Body = []string{"user already exists"}
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(errorResponse)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	apiUser := converter.ToAPIUser(u)

	resp := generated.UserResponse{
		User: *apiUser,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req generated.LoginUser
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	mapReq := user.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	}
	u, err := h.AuthService.Login(mapReq)

	if err != nil {
		switch err {
		case user.ErrInvalidCredentials:
			errorResponse := generated.GenericErrorModel{}
			errorResponse.Errors.Body = []string{"invalid credentials"}
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(errorResponse)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	apiUser := converter.ToAPIUser(u)

	resp := generated.UserResponse{
		User: *apiUser,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
