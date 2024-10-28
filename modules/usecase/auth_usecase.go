package usecase

import (
	"go-clean-arch/modules/entities/auth"
	"go-clean-arch/modules/repositories/authrepo"

	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase struct {
	repo authrepo.AuthRepository
}

func NewAuthUsecase(repo authrepo.AuthRepository) *AuthUsecase {
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
	err := uc.repo.HandleSignUp(detail)
	return err
}

func (uc *AuthUsecase) HandleSignIn(username string) (auth.UserAuth, error) {
	userDetail, err := uc.repo.HandleSignIn(username)
	return userDetail, err
}
