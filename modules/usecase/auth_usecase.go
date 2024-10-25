package usecase

import (
	"go-clean-arch/modules/entities/auth"
	"go-clean-arch/modules/repositories"

	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase struct {
	repo repositories.AuthRepository
}

func NewAuthUsecase(repo repositories.AuthRepository) *AuthUsecase {
	return &AuthUsecase{repo: repo}
}

func (uc *AuthUsecase) IsUsernameExist(username string) bool {
	isExist := uc.repo.IsUsernameExist(username)
	return isExist
}

func (uc *AuthUsecase) HashPassword(password string) (string, error) {
	// 8 is cost for generate hash bigger is mean slower
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	return string(bytes), err
}

func (uc *AuthUsecase) HandleSignUp(detail auth.UserAuth) error {
	error := uc.repo.HandleSignUp(detail)
	return error
}
