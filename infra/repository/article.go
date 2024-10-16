package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"

	"handler/function/domain/model"
)

const ArticleCollection = "article"

func (r *Repository) GetArticles(ctx context.Context) ([]*model.Article, error) {
	repositoryEB := repositoryEB.Tags("GetArticles")

	cursor, err := r.db.Collection(ArticleCollection).Find(ctx, bson.D{})
	if err != nil {
		return nil, repositoryEB.
			Tags("db.Collection.Find").
			Wrap(err)
	}

	var articles []*model.Article
	if err := cursor.All(ctx, &articles); err != nil {
		return nil, repositoryEB.
			Tags("cursor.All").
			Wrap(err)
	}

	return articles, nil
}

func (r *Repository) DeleteArticle(ctx context.Context, article *model.Article) error {
	deleteResult, err := r.db.Collection(ArticleCollection).DeleteOne(ctx, bson.M{"_id": article.ID})
	if err != nil {
		return repositoryEB.
			Tags("db.Collection.DeleteOne").
			Wrap(err)
	}

	if deleteResult.DeletedCount == 0 {
		return repositoryEB.
			Tags("DeletedCount").
			Errorf("No article deleted.")
	}

	return nil
}
