package main

import (
	"fmt"

	"github.com/Ara-Chobanyan/student-average/pkg/average"
	"github.com/Ara-Chobanyan/student-average/pkg/students"
)

func main() {
	a := average.FindAverage([]int{1, 2, 3, 4, 5})
	fmt.Println(a)
	b := students.GetUserInput()
}
