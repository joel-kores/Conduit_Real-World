package handlers

import (
    "net/http"

    "github.com/joel-kores/Conduit_Real-World/pkg/logger"
)

type TagHandler struct {
    Logger logger.Logger
}

func NewTagHandler(logger logger.Logger) *TagHandler {
    return &TagHandler{
        Logger: logger,
    }
}

func (h *TagHandler) GetTags(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusNotImplemented)
    w.Write([]byte(`{"errors":{"body":["GetTags not implemented yet"]}}`))
}
