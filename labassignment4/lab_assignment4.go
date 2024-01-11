package main

import (
	"fmt"
	"strings"
)

type student struct {
	student_name string
	student_id   int
	grades       [4]int
}

func main() {
	displayPrompt()
}
func displayPrompt() {
	fmt.Println("This program accepts 4 test grades for the student and it calculates the average of the four highest grades.")
	fmt.Println("Enter student's name:")
	var name string
	fmt.Scan(&name)
	fmt.Println("Enter ID number:")
	var idNumber int
	fmt.Scan(&idNumber)

	count := 0
	var grade int
	var gradeArray []int
	// fmt.Printf("len=%d cap=%d %v\n", len(gradeArray), cap(gradeArray), gradeArray)
	for count < 4 {
		fmt.Printf("Enter grade for test %d: ", count+1)
		fmt.Scan(&grade)
		gradeArray = append(gradeArray, grade)
		count++
	}
	// fmt.Printf("len=%d cap=%d %v\n", len(gradeArray), cap(gradeArray), gradeArray)
	/*
		for index, data := range gradeArray {
			fmt.Println(index)
			fmt.Println(data)
		}
	*/
	studentData := student{student_name: name, student_id: idNumber, grades: [4]int(gradeArray)}

	var average float32 = calculateAverage(studentData)

	displayData(studentData, average)
}
func calculateAverage(studentData student) float32 {
	var sum int = 0
	for _, data := range studentData.grades {
		sum += data
	}
	return float32(sum) / float32(len(studentData.grades))
}
func displayData(studentData student, average float32) {
	fmt.Printf("Name.............................%s \n", studentData.student_name)
	fmt.Printf("Id Number.........................%d \n", studentData.student_id)
	fmt.Printf("Test grades.........................%d, %d, %d, %d \n", studentData.grades[0], studentData.grades[1], studentData.grades[2], studentData.grades[3])
	fmt.Printf("Avearage of the four highest grades%.2f \n", average)
	fmt.Println("------------------------------------")
	fmt.Println("Do you want to quit")
	fmt.Println("Enter `Q` to quit or any other character to continue.")
	var usercontinue string
	fmt.Scan(&usercontinue)
	if strings.ToLower(usercontinue) != "q" {
		main()
	}
}
