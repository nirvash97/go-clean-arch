package usecase

import (
	"go-clean-arch/modules/entities/movies"
	repositories "go-clean-arch/modules/repositories/movie"
)

type MovieUsecase struct {
	repo repositories.MovieRepository
}

func NewMovieUsecase(repo repositories.MovieRepository) *MovieUsecase {
	return &MovieUsecase{repo: repo}
}

func (uc *MovieUsecase) GetMoviesTest() ([]movies.Movie, error) {
	return uc.repo.GetMoviesTest()
}

func (uc *MovieUsecase) GetMovieByLanguage(language string) ([]movies.Movie, error) {
	movie, err := uc.repo.GetMovieByLanguage(language)
	return movie, err
}

func (uc *MovieUsecase) GetMovieByLanguagePagination(language string, page int64, perPage int64) ([]movies.Movie, error) {
	movie, err := uc.repo.GetMovieByLanguagePagination(language, page, perPage)
	return movie, err
}

func (uc *MovieUsecase) GetMovieByLanguageItemCount(language string) (int64, error) {
	itemCount, err := uc.repo.GetMovieByLanguageItemCount(language)
	return itemCount, err
}
