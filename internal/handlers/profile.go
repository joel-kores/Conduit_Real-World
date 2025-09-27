package handlers

import (
	"net/http"

	"github.com/joel-kores/Conduit_Real-World/pkg/logger"
)

type ProfileHandler struct {
	Logger logger.Logger
}

func NewProfileHandler(logger logger.Logger) *ProfileHandler {
	return &ProfileHandler{
		Logger: logger,
	}
}

func (h *ProfileHandler) GetProfileByUsername(w http.ResponseWriter, r *http.Request, username string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte(`{"errors":{"body":["GetProfileByUsername not implemented yet"]}}`))
}

func (h *ProfileHandler) FollowUserByUsername(w http.ResponseWriter, r *http.Request, username string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte(`{"errors":{"body":["FollowUserByUsername not implemented yet"]}}`))
}

func (h *ProfileHandler) UnfollowUserByUsername(w http.ResponseWriter, r *http.Request, username string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte(`{"errors":{"body":["UnfollowUserByUsername not implemented yet"]}}`))
}
