package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/Ara-Chobanyan/student-average/cmd/average"
	"github.com/Ara-Chobanyan/student-average/cmd/students"
)

func main() {
	//Create a map to store new students
	b := make(map[string]float32)
	fmt.Println("Enter the amount of students")
	//Enter the amount of students you wish to make

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	//Converts the user input into a string
	name := scanner.Text()
	num, _ := strconv.Atoi(name)

	for i := 0; i < num; i++ {
		createStudent := students.GetStudent()
		// createStudent.OutPutStudents()
		a := average.FindAverage(createStudent.Grades)
		b[createStudent.Name] = a
	}
	fmt.Println(b)
}
