package main

import (
	"StudentGradeBookManager/student"
	"fmt"
)

func main() {

	fmt.Println("------ Student Grade Book--------- \n")

	alice := student.CreateStudent("Alice", 20, 95)
	bob := student.CreateStudent("Bob", 22, 85)
	charlie := student.CreateStudent("Charlie", 21, 75)

	students := []student.Student{alice, bob, charlie}

	fmt.Println("-------Student List -------")
	for _, s := range students {
		s.PrintDetails()
		fmt.Println()
	}

	fmt.Println("-------Top Student -------")
	student.CalculateTopStudent(students).PrintDetails()

	fmt.Println()

	fmt.Println("Average Score of students is ", student.CalculateAverageMarks(students))
}

/* AP3X. */
