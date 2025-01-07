package examrepo

import (
	"context"
	"go-clean-arch/modules/entities/exam"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ExamMongoRepo struct {
	collection *mongo.Collection
}

func NewExamMongoRepo(client *mongo.Client) *ExamMongoRepo {
	db := client.Database("exam_mongo")
	return &ExamMongoRepo{
		collection: db.Collection("exam_user"),
	}
}

func (r *ExamMongoRepo) GetAllUser() ([]exam.ExamUser, error) {
	var examUser []exam.ExamUser
	option := options.Find()
	cursor, err := r.collection.Find(context.Background(), bson.D{}, option)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	if err = cursor.All(context.Background(), &examUser); err != nil {
		return nil, err
	}
	return examUser, err
}

func (r *ExamMongoRepo) PostAddUser(name string, email string) error {

	_, err := r.collection.InsertOne(context.Background(), bson.D{{Key: "name", Value: name}, {Key: "email", Value: email}})
	if err != nil {
		return err
	}
	return nil
}

func (r *ExamMongoRepo) GetUserById(id int) (*exam.ExamUser, error) {
	var userDetail exam.ExamUser

	filter := bson.D{{Key: "id", Value: id}}

	err := r.collection.FindOne(context.Background(), filter).Decode(&userDetail)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &userDetail, nil
}

func (r *ExamMongoRepo) PutUpdateUser(id int, name string, email string) (*exam.ExamUser, error) {
	var userDetail exam.ExamUser
	findFilter := bson.D{{Key: "id", Value: id}}
	updateData := bson.D{{Key: "name", Value: name}, {Key: "email", Value: email}}
	updateFilter := bson.D{{Key: "$set", Value: updateData}}

	_, err := r.collection.UpdateOne(context.Background(), findFilter, updateFilter)
	if err != nil {
		return nil, err
	}
	findErr := r.collection.FindOne(context.Background(), findFilter).Decode(&userDetail)
	if findErr != nil {
		return nil, err
	}
	return &userDetail, nil
}
