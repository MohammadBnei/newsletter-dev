package router

import (
	"context"

	"github.com/go-chi/chi/v5"

	"handler/function/domain/model"
)

type Service interface {
	GetArticles(ctx context.Context) ([]*model.Article, error)
	DeleteArticle(ctx context.Context, article *model.Article) error
}

type Router struct {
	chi.Router
	service Service
}

func NewRouter(service Service) *Router {
	return &Router{
		service: service,
	}
}
