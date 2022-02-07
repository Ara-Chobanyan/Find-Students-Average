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
	b := make(map[string]float32)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
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
