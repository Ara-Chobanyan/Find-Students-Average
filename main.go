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
	amount := scanner.Text()
	num, err := strconv.Atoi(amount)
	if err != nil {
		fmt.Println("Please enter a valid number")
		return
	}

	for i := 0; i < num; i++ {
		createStudent := students.GetStudent()
		// createStudent.OutPutStudents()
		a := average.FindAverage(createStudent.Grades)
		b[createStudent.Name] = a
	}
	fmt.Println(b)
}

/*
I can take this a step futher and give options to put in the students grade.
Then it reset back at the menu and you can recheck the grades
*/
