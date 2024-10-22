package repositories

import "go-clean-arch/modules/entities/movies"

type MovieRepository interface {
	GetMovieByLanguage(languages string) ([]movies.Movie, error)
	//GetAllMovie() ([]movies.Movie, error)
}
