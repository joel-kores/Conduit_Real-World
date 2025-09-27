package handlers

import (
	"net/http"

	"github.com/joel-kores/Conduit_Real-World/pkg/logger"
)

type CommentHandler struct{
	logger logger.Logger
}

func NewCommentHandler(logger logger.Logger) *CommentHandler {
	return &CommentHandler{
		logger: logger,
	}
}

func (h *CommentHandler) GetArticleComments(w http.ResponseWriter, r *http.Request, slug string) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusNotImplemented)
    w.Write([]byte(`{"errors":{"body":["GetArticleComments not implemented yet"]}}`))
}

func (h *CommentHandler) CreateArticleComment(w http.ResponseWriter, r *http.Request, slug string) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusNotImplemented)
    w.Write([]byte(`{"errors":{"body":["CreateArticleComment not implemented yet"]}}`))
}

func (h *CommentHandler) DeleteArticleComment(w http.ResponseWriter, r *http.Request, slug string, id int) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusNotImplemented)
    w.Write([]byte(`{"errors":{"body":["DeleteArticleComment not implemented yet"]}}`))
}
