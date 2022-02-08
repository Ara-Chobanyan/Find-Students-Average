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

// func (c *Student) OutPutStudents() {
// 	fmt.Println(c.Name, c.Grades)
// }

// NewStudent - takes in a string and a []int then returns a pointer to Student to avoid duplications
func NewStudent(name string, grades []int) *Student {
	return &Student{name, grades}
}

// GetStudent - Starts the functions that create new students and returns them
func GetStudent() *Student {
	student := getAdminInput()
	return student
}

//getAdminInput - Takes in the users inputs and turns them to a Student Data type
func getAdminInput() *Student {
	var name string
	var grades []int
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Please Enter Name")
	scanner.Scan()
	name = scanner.Text()
	//Bad code here repeating my self need to refactor this
	fmt.Println("Please Enter Grades with space in between")
	scanner.Scan()
	grades = convertStringToSlice(scanner.Text())

	student := NewStudent(name, grades)
	return student
}

// convertStringToSlice - Converts strings of numbers into []int
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
