package examrepo

import "go-clean-arch/modules/entities/exam"

type ExamRepository interface {
	GetAllUser() ([]exam.ExamUser, error)
	PostAddUser(name string, email string) error
	GetUserById(id int) (*exam.ExamUser, error)
	PutUpdateUser(id int, name string, email string) (*exam.ExamUser, error)
}
