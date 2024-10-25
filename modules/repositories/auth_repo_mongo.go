package repositories

import (
	"context"
	"go-clean-arch/modules/entities/auth"

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
