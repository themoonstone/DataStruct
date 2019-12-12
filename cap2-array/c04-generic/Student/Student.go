package Student

import "fmt"

type Student struct {
	Name  string
	Score int
}

func (s *Student) String() string {
	return fmt.Sprintf("Student(name:%s, score:%d)", s.Name, s.Score)
}
