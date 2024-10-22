package movies

import (
	"time"
)

type Movie struct {
	Title           string    `bson:"title"`
	Genres          []string  `bson:"genres"`
	Runtime         int       `bson:"runtime"`
	Imdb            MovieImdb `bson:"imdb"`
	Plot            string    `bson:"plot"`
	Rated           string    `bson:"rated"`
	Cast            []string  `bson:"cast"`
	NumMflixComment int       `bson:"num_mflix_comments"`
	Poster          string    `bson:"poster"`
	Fullplot        string    `bson:"fullplot"`
	Languages       []string  `bson:"languages"`
	Released        time.Time `bson:"released"`
	Directors       []string  `bson:"directors"`
	Writers         []string  `bson:"writers"`
}

type MovieImdb struct {
	Rating float64 `bson:"rating"`
	Voting int64   `bson:"votes"`
	Dd     int64   `bson:"id"`
}
