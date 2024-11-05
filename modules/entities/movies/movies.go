package movies

import (
	"time"
)

// Tips
// 1.) omitempty will hide struct field when value was nil
// 2.) prefix struct variable with * mean it be nullable ( back to read first tips )
type Movie struct {
	Title           string     `json:"title"`
	Genres          []string   `json:"genres,omitempty"`
	Year            int        `json:"year"`
	Runtime         int        `json:"runtime,omitempty"`
	Imdb            *MovieImdb `json:"imdb,omitempty"`
	Plot            string     `json:"plot,omitempty"`
	Rated           string     `json:"rated,omitempty"`
	Cast            []string   `json:"cast,omitempty"`
	NumMflixComment int        `json:"num_mflix_comments,omitempty"`
	Poster          string     `json:"poster,omitempty"`
	Fullplot        string     `json:"fullplot,omitempty"`
	Languages       []string   `json:"languages,omitempty"`
	Released        *time.Time `json:"released,omitempty"`
	Directors       []string   `json:"directors,omitempty"`
	Writers         []string   `json:"writers,omitempty"`
}

type MovieImdb struct {
	Rating float64 `json:"rating,omitempty"`
	Voting int64   `json:"votes,omitempty"`
	Dd     int64   `json:"id,omitempty"`
}
