package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/Ara-Chobanyan/find-students-average/logic/average"
	"github.com/Ara-Chobanyan/find-students-average/logic/menu"
	"github.com/Ara-Chobanyan/find-students-average/logic/students"
)

func main() {
	//Create a map to store new students
	b := make(map[string]float32)

	for {
		// Begins the program
		a, err := menu.MenuOptions()
		if err != nil {
			fmt.Println(err)
		}

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
			// Deletes a student from the map
			fmt.Println("Enter Student name")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			scanner.Text()
			//Takes in the key which is the students name and deletes it

			name := scanner.Text()
			delete(b, name)
		} else if a == "4" {
			// Updates a student from the map
			fmt.Println("Enter Student name")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			scanner.Text()
			name := scanner.Text()

			var f float32
			fmt.Println("Enter A New Grade")
			_, err := fmt.Scanf("%f", &f)

			if err != nil {
				fmt.Println(err)
			}

			b[name] = f
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
