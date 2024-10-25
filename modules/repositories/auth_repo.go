package repositories

import "go-clean-arch/modules/entities/auth"

type AuthRepository interface {
	IsUsernameExist(username string) bool
	HandleSignUp(detail auth.UserAuth) error
}
