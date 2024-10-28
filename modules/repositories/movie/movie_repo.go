package movierepo

import (
	"go-clean-arch/modules/entities/movies"
)

type MovieRepository interface {
	GetMovieByLanguage(languages string) ([]movies.Movie, error)
	//GetAllMovie() ([]movies.Movie, error)
	GetMovieByLanguagePagination(language string, page int64, perPage int64) ([]movies.Movie, error)
	GetMovieByLanguageItemCount(language string) (int64, error)
}
