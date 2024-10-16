package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Article struct {
	ID      primitive.ObjectID `bson:"_id"`
	Link    string
	Content string
	Week    string
}
