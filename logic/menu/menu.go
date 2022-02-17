package menu

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

// MenuOptions - Starts the menu for the user to pick there options
func MenuOptions() (string, error) {
	var choice string

	fmt.Println("Please make your choice")
	fmt.Println("1) Create students")
	fmt.Println("2) Read a already made group of students")
	fmt.Println("3) Delete a Student from the group")
	fmt.Println("4) Exit")

	// Takes in the user input
	userInput := bufio.NewScanner(os.Stdin)
	userInput.Scan()

	userChoice := userInput.Text()
	// Assign it to choice to see what option the user picked
	choice = userChoice

	if choice == "1" || choice == "2" || choice == "3" || choice == "4" {
		return choice, nil
	} else {
		return "", errors.New("Invalid input")
	}
}
