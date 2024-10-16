package repository

import (
	"github.com/samber/oops"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	db *mongo.Database
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		db: db,
	}
}

var repositoryEB = oops.In("repository").With("dbType", "mongo")
