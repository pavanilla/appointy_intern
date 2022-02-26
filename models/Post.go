package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	UniqueID  primitive.ObjectID `json:"_id" bson"_id"`
	Author    primitive.ObjectID `json:"authorId" bson:"authorId"`
	Posted    primitive.DateTime `json:"posted" bson:"posted`
	Title     string             `json:"title" bson:"tile"`
	Body      string             `json:"body"  bson:"body"`
	Thumbnail string             `json:"thumbnail" bson:"thumbnail,omitempty`
}
