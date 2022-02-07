package students

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	Name   string
	Grades []int
}

func (c *Student) OutPutStudents() {
	fmt.Println(c.Name, c.Grades)
}

func NewStudent(name string, grades []int) *Student {
	return &Student{name, grades}
}

//Start of creaing students
func GetStudent() *Student {
	student := getAdminInput()
	return student
}

func getAdminInput() *Student {
	var name string
	var grades []int
	scanner := bufio.NewScanner(os.Stdin)
	for i := 1; i <= 2 && scanner.Scan(); i++ {
		switch i {
		case 1:
			name = scanner.Text()
		case 2:
			grades = convertStringToSlice(scanner.Text())
		}
	}
	student := NewStudent(name, grades)
	return student
}

func convertStringToSlice(s string) []int {
	var n []int
	for _, f := range strings.Fields(s) {
		i, err := strconv.Atoi(f)
		if err == nil {
			n = append(n, i)
		}
	}
	return n
}
