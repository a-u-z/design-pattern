package factory

import "context"

type StudentDaoType int

type Student struct {
	ID   int
	Name string
}

type StudentDao interface {
	Create(ctx context.Context, stu Student) bool
}

const (
	StudentDaoA StudentDaoType = 1
	StudentDaoB StudentDaoType = 2
)

type StudentDaoImp1 struct{}

func NewStudentDaoImp1() *StudentDaoImp1 {
	return &StudentDaoImp1{}
}
