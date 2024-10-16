package service

import (
	"context"

	"handler/function/domain/model"
)

func (s *Service) GetArticles(ctx context.Context) ([]*model.Article, error) {
	serviceEB := serviceEB.Tags("GetArticles")

	articles, err := s.repository.GetArticles(ctx)
	if err != nil {
		return nil, serviceEB.Wrap(err)
	}

	return articles, nil
}

func (s *Service) DeleteArticle(ctx context.Context, article *model.Article) error {
	serviceEB := serviceEB.Tags("DeleteArticle")

	err := s.repository.DeleteArticle(ctx, article)
	if err != nil {
		return serviceEB.Wrap(err)
	}

	return nil
}
