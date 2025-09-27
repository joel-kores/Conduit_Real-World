package handlers

import (
	"github.com/joel-kores/Conduit_Real-World/internal/services"
	"github.com/joel-kores/Conduit_Real-World/pkg/logger"
)

// CompleteHandler embeds all handlers to satisfy ServerInterface
type CompleteHandler struct {
	*AuthHandler
	*ArticleHandler
	*CommentHandler
	*ProfileHandler
	*TagHandler
}

func NewCompleteHandler(authService *services.AuthService, logger logger.Logger) *CompleteHandler {
	return &CompleteHandler{
		AuthHandler:    NewAuthHandler(authService, logger),
		ArticleHandler: NewArticleHandler(logger),
		CommentHandler: NewCommentHandler(logger),
		ProfileHandler: NewProfileHandler(logger),
		TagHandler:     NewTagHandler(logger),
	}
}
