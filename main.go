package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/Ara-Chobanyan/student-average/cmd/average"
	"github.com/Ara-Chobanyan/student-average/cmd/menu"
	"github.com/Ara-Chobanyan/student-average/cmd/students"
)

func main() {
	//Create a map to store new students
	b := make(map[string]float32)

	for {
		a, _ := menu.MenuOptions()

		if a == "1" {
			fmt.Println("Enter the amount of students")
			//Enter the amount of students you wish to make
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()

			//Converts the user input into a int
			amount := scanner.Text()
			num, err := strconv.Atoi(amount)
			if err != nil {
				fmt.Println("Please enter a valid number")
				return
			}
			amountOfStudents(num, b)
		} else if a == "2" {
			fmt.Println("---------------------")
			fmt.Println("Group:", b)
			fmt.Println("---------------------")
		} else if a == "3" {
			fmt.Println("Enter Student name")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			scanner.Text()

			name := scanner.Text()
			delete(b, name)
		} else {
			break
		}
	}

}

// amountOfStudents - Takes in a int and a map to create a map of students with there grades and find the students average
func amountOfStudents(num int, group map[string]float32) {
	for i := 0; i < num; i++ {
		createStudent := students.GetStudent()

		a := average.FindAverage(createStudent.Grades)
		group[createStudent.Name] = a
	}
}
