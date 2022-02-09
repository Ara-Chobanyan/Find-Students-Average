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
			fmt.Println(b)
		} else {
			break
		}
	}

}

func amountOfStudents(num int, group map[string]float32) {
	for i := 0; i < num; i++ {
		createStudent := students.GetStudent()

		a := average.FindAverage(createStudent.Grades)
		group[createStudent.Name] = a
	}
}

/*
I can take this a step futher and give options to put in the students grade.
Then it reset back at the menu and you can recheck the grades


Create a infinite for loop that gives the user 3 options
1. Create a new list of students.
2. Look at the current list of students will be empty at first.
3. Quit

Need
*/
