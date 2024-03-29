package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRequestModel struct {
	FirstName string `bson:"first_name" json:"first_name"`
	LastName  string `bson:"last_name" json:"last_name"`
	Email     string `bson:"email" json:"email"`
	Username  string `bson:"username" json:"username"`
	Password  string `bson:"password" json:"password"`
}

type UserDbModel struct {
	*UserRequestModel
	Id        primitive.ObjectID `bson:"_id" json:"id"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

type UserDb interface {
	CreateUser(user *UserRequestModel) (bool, error)
}
