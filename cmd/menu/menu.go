package menu

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func MenuOptions() (string, error) {
	var choice string

	fmt.Println("Please make your choice")
	fmt.Println("1) Create a new group of students")
	fmt.Println("2) Read a already made group of students")
	fmt.Println("3) Exit")

	userInput := bufio.NewScanner(os.Stdin)
	userInput.Scan()

	userChoice := userInput.Text()
	choice = userChoice

	if choice == "1" || choice == "2" || choice == "3" {
		return choice, nil
	} else {
		return "", errors.New("Invalid input")
	}
}

// func CreateStudents() {
// 	for {
// 		startMenu, err := MenuOptions()
// 		if err != nil {
// 			fmt.Println("Invalid option")
// 			return
// 		}
//
// 		if startMenu == "1" {
// 			//do something
// 		} else if startMenu == "2" {
// 			//do something
// 		} else {
// 			return
// 		}
// 	}
// }

/*
GOAL: Create a simple menu type of function to allow the user to pick a option
The options should be make a new group of students,Read the group of students or to quit
*/
