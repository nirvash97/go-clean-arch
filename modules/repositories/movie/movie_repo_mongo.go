package movierepo

import (
	"context"
	"go-clean-arch/modules/entities/movies"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MovieMongoRepo struct {
	collection *mongo.Collection
}

func NewMovieMongoRepo(client *mongo.Client) *MovieMongoRepo {
	db := client.Database("sample_mflix")
	return &MovieMongoRepo{
		collection: db.Collection("movies"),
	}
}

func (r *MovieMongoRepo) GetMoviesTest() ([]movies.Movie, error) {
	var movies []movies.Movie
	option := options.Find()
	option.SetLimit(5)
	cursor, err := r.collection.Find(context.Background(), bson.D{}, option)
	if err != nil {
		println("Error Collection ")

		return nil, err
	}
	defer cursor.Close(context.Background())
	if err = cursor.All(context.Background(), &movies); err != nil {
		println("Error Cursor ")
		return nil, err
	}
	return movies, err
}

func (r *MovieMongoRepo) GetMovieByLanguage(language string) ([]movies.Movie, error) {
	var movies []movies.Movie
	filter := bson.D{{Key: "languages", Value: bson.D{{Key: "$in", Value: bson.A{language}}}}}
	option := options.Find()
	option.SetLimit(10)
	// option.SetProjection(bson.D{
	// 	{Key: "title", Value: 1},
	// 	{Key: "year", Value: 1},
	// 	{Key: "runtime", Value: 1},
	// 	{Key: "imdb", Value: 1},
	// 	// {Key: "released", Value: 1},
	// 	{Key: "languages", Value: 1},
	// })
	cursor, err := r.collection.Find(context.Background(), filter, option)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	if err = cursor.All(context.Background(), &movies); err != nil {
		return nil, err
	}

	return movies, nil

}

func (r *MovieMongoRepo) GetMovieByLanguagePagination(language string, page int64, perPage int64) ([]movies.Movie, error) {
	var movies []movies.Movie
	filter := bson.D{{Key: "languages", Value: bson.D{{Key: "$in", Value: bson.A{language}}}}}
	option := options.Find()
	option.SetLimit(10)
	skip := (page - 1) * perPage
	option.SetSkip(skip)

	cursor, err := r.collection.Find(context.Background(), filter, option)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	if err = cursor.All(context.Background(), &movies); err != nil {
		return nil, err
	}

	return movies, nil
}

func (r *MovieMongoRepo) GetMovieByLanguageItemCount(language string) (int64, error) {

	filter := bson.D{{Key: "languages", Value: bson.D{{Key: "$in", Value: bson.A{language}}}}}

	itemCount, err := r.collection.CountDocuments(context.Background(), filter)
	if err != nil {
		return 0, err
	}

	return itemCount, nil
}
