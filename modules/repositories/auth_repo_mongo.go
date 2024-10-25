package repositories

import (
	"context"
	"go-clean-arch/modules/entities/auth"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthMongoRepo struct {
	collection *mongo.Collection
}

func NewAuthMongoRepo(db *mongo.Database) *AuthMongoRepo {
	return &AuthMongoRepo{
		collection: db.Collection("auth"),
	}
}

func (r *AuthMongoRepo) IsUsernameExist(username string) bool {
	filter := bson.D{{Key: "username", Value: username}}
	var existUser auth.UserAuth
	err := r.collection.FindOne(context.Background(), filter).Decode(&existUser)
	if err == nil {
		log.Println(existUser)
		return true
	} else {
		return false
	}

}

func (r *AuthMongoRepo) HandleSignUp(detail auth.UserAuth) error {
	_, err := r.collection.InsertOne(context.Background(), detail)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (r *AuthMongoRepo) HandleSignIn(username string) (auth.UserAuth, error) {
	var user auth.UserAuth
	filter := bson.D{{Key: "username", Value: username}}
	err := r.collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return auth.UserAuth{}, err
	} else {
		return user, nil
	}
}
