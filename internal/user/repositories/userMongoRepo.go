package repositories

import (
	"context"
	"go_api/world/internal/user/domain"
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

func (ur *UserMongoRepository) CreateUser(user []byte) (bool, error) {

	result, mongoErr := ur.collection.InsertOne(ur.ctx, user)

	if mongoErr != nil {
		log.Fatal(mongoErr)
	}

	log.Println(result)

	return true, nil
}

func (ur *UserMongoRepository) FindUserByUserName(username string) (*domain.User, error) {

	var userResult domain.User

	err := ur.collection.FindOne(ur.ctx, bson.D{{Key: "username", Value: username}}).Decode(&userResult)

	return &userResult, err
}

func (ur *UserMongoRepository) FindUserByEmail(email string) (*domain.User, error) {

	var userResult domain.User

	err := ur.collection.FindOne(ur.ctx, bson.D{{Key: "email", Value: email}}).Decode(&userResult)

	return &userResult, err
}
