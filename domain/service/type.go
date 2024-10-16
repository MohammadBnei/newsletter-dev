package service

import (
	"context"

	"github.com/samber/oops"

	"handler/function/domain/model"
)

type Service struct {
	repository serviceInterface
}

type serviceInterface interface {
	GetArticles(ctx context.Context) ([]*model.Article, error)
	DeleteArticle(ctx context.Context, article *model.Article) error
}

func NewService(repository serviceInterface) *Service {
	return &Service{
		repository,
	}
}

var serviceEB = oops.In("service")
