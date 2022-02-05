package students

import (
	"fmt"
)

type Student struct {
	Name   string
	Grades []int
}

func (s Student) OutputStudents() {
	fmt.Println(s.Name, s.Grades)
}

  func ()  {
    
  }
