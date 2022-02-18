package students

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Ara-Chobanyan/find-students-average/logic/stringtoslice"
)

type Student struct {
	Name   string
	Grades []int
}

// NewStudent - takes in a string and a []int then returns a pointer to Student to avoid duplications
func NewStudent(name string, grades []int) *Student {
	return &Student{name, grades}
}

// GetStudent - Starts the functions that create new students and returns them
func GetStudent() *Student {
	student := getAdminInput()
	return student
}

// getAdminInput - Takes in the users inputs and turns them to a Student Data type
func getAdminInput() *Student {
	var name string
	var grades []int
	var student *Student

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Please Enter Name")
	scanner.Scan()
	name = scanner.Text()

	fmt.Println("Please Enter Grades with space in between")
	scanner.Scan()
	grades = stringtoslice.ConvertStringToSlice(scanner.Text())

	student = NewStudent(name, grades)
	return student
}
