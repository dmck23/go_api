package repositories

import (
	"context"
	"go_api/world/internal/user/domain"
	"go_api/world/internal/user/errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserMongoRepository struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewUserMongoRepository() domain.UserDb {

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	userCollection := client.Database("Users").Collection("users")

	return &UserMongoRepository{collection: userCollection, ctx: context.TODO()}
}

func (ur *UserMongoRepository) CreateUser(user *domain.UserRequestModel) (bool, error) {

	// var db_user domain.UserDbModel

	var userExistsResult bson.M

	userErr := ur.collection.FindOne(ur.ctx, bson.D{{Key: "username", Value: user.Username}}).Decode(&userExistsResult)

	emailErr := ur.collection.FindOne(ur.ctx, bson.D{{Key: "email", Value: user.Email}}).Decode(&userExistsResult)

	if userErr == mongo.ErrNoDocuments && emailErr == mongo.ErrNoDocuments {
		_, mongoErr := ur.collection.InsertOne(ur.ctx, user)

		if mongoErr != nil {
			log.Fatal(mongoErr)
		}
		return true, nil
	}

	return false, errors.UserAlreadyExists{}
}
