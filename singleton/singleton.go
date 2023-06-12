package singleton

import "sync"

type Student struct {
	ID   int
	Name string
}

var (
	defaultStudent *Student
)

func NewStudent() *Student {
	once := sync.Once{}

	once.Do(func() {
		defaultStudent = &Student{}
	})
	return defaultStudent
}
