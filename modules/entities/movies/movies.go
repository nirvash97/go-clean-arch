package movies

import (
	"time"
)

type Movie struct {
	Title           string    `json:"title"`
	Genres          []string  `json:"genres"`
	Runtime         int       `json:"runtime"`
	Imdb            MovieImdb `json:"imdb"`
	Plot            string    `json:"plot"`
	Rated           string    `json:"rated"`
	Cast            []string  `json:"cast"`
	NumMflixComment int       `json:"num_mflix_comments"`
	Poster          string    `json:"poster"`
	Fullplot        string    `json:"fullplot"`
	Languages       []string  `json:"languages"`
	Released        time.Time `json:"released"`
	Directors       []string  `json:"directors"`
	Writers         []string  `json:"writers"`
}

type MovieImdb struct {
	Rating float64 `bson:"rating"`
	Voting int64   `bson:"votes"`
	Dd     int64   `bson:"id"`
}
