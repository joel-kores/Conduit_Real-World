package handlers

import (
	"net/http"

	"github.com/joel-kores/Conduit_Real-World/internal/generated"
	"github.com/joel-kores/Conduit_Real-World/pkg/logger"
)

type ArticleHandler struct {
	Logger logger.Logger
}

func NewArticleHandler(logger logger.Logger) *ArticleHandler {
	return &ArticleHandler{
		Logger: logger,
	}
}

func (h *ArticleHandler) GetArticles(w http.ResponseWriter, r *http.Request, params generated.GetArticlesParams) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte(`{"errors":{"body":["GetArticles not implemented yet"]}}`))
}

func (h *ArticleHandler) CreateArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte(`{"errors":{"body":["CreateArticle not implemented yet"]}}`))
}

func (h *ArticleHandler) GetArticle(w http.ResponseWriter, r *http.Request, slug string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte(`{"errors":{"body":["GetArticle not implemented yet"]}}`))
}

func (h *ArticleHandler) UpdateArticle(w http.ResponseWriter, r *http.Request, slug string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte(`{"errors":{"body":["UpdateArticle not implemented yet"]}}`))
}

func (h *ArticleHandler) DeleteArticle(w http.ResponseWriter, r *http.Request, slug string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte(`{"errors":{"body":["DeleteArticle not implemented yet"]}}`))
}

func (h *ArticleHandler) CreateArticleFavorite(w http.ResponseWriter, r *http.Request, slug string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte(`{"errors":{"body":["CreateArticleFavorite not implemented yet"]}}`))
}

func (h *ArticleHandler) DeleteArticleFavorite(w http.ResponseWriter, r *http.Request, slug string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte(`{"errors":{"body":["DeleteArticleFavorite not implemented yet"]}}`))
}

func (h *ArticleHandler) GetArticlesFeed(w http.ResponseWriter, r *http.Request, params generated.GetArticlesFeedParams) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte(`{"errors":{"body":["GetArticlesFeed not implemented yet"]}}`))
}
