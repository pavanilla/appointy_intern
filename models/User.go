package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	UniqueID    primitive.ObjectID `json:"_id" bson:"_id`
	Name        string             `json:"name" bson:"name"`
	Email       string             `json:"email"  bson:"email"`
	Username    string             `json:"username unique" bson:"username"`
	PhoneNumber string             `json:"phonenumber" bson:"phonenumber,omitempty"`
	Password    string             `json:"password" bson:"password omitempty"`
	Dob         primitive.DateTime `json:"date" bson:"date,omitempty"`
}
