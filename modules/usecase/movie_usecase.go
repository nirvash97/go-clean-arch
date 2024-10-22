package usecase

import (
	"go-clean-arch/modules/entities/movies"
	"go-clean-arch/modules/repositories"
)

type MovieUsecase struct {
	repo repositories.MovieRepository
}

func NewMovieUsecase(repo repositories.MovieRepository) *MovieUsecase {
	return &MovieUsecase{repo: repo}
}

func (uc *MovieUsecase) GetMovieByLanguage(language string) ([]movies.Movie, error) {
	movie, err := uc.repo.GetMovieByLanguage(language)
	return movie, err
}
