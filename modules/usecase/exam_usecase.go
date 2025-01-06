package usecase

import (
	"go-clean-arch/modules/entities/exam"
	repositories "go-clean-arch/modules/repositories/examerpo"
)

type ExamUseCase struct {
	repo repositories.ExamRepository
}

func NewExamUsecase(repo repositories.ExamRepository) *ExamUseCase {
	return &ExamUseCase{repo: repo}
}

func (uc *ExamUseCase) ExamHelloWorldText() string {
	return "Hello World GO"
}

func (uc *ExamUseCase) GetAllUser() ([]exam.ExamUser, error) {
	userList, err := uc.repo.GetAllUser()
	return userList, err
}

func (uc *ExamUseCase) PostAdduser(name string, email string) error {
	err := uc.repo.PostAddUser(name, email)
	return err
}

func (uc *ExamUseCase) GetUserById(id int) (*exam.ExamUser, error) {
	userDetail, err := uc.repo.GetUserById(id)
	return userDetail, err
}
func (uc *ExamUseCase) PutUpdateUser(id int, name string, email string) (*exam.ExamUser, error) {
	userDetail, err := uc.repo.PutUpdateUser(id, name, email)
	return userDetail, err
}
