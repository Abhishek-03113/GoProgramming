package student

import "fmt"

type Student struct {
	Name  string
	Age   int
	Marks int
	Grade string
}

func (s Student) PrintDetails() {
	fmt.Printf("Name: %s\n | Age: %d\n | Marks: %d\n | Grade: %s\n", s.Name, s.Age, s.Marks, s.Grade)
}

func CreateStudent(name string, age int, marks int) Student {
	grade := calculateGrade(marks)
	return Student{Name: name, Age: age, Marks: marks, Grade: grade}
}

func calculateGrade(marks int) string {
	if marks >= 90 {
		return "A"
	} else if marks >= 80 {
		return "B"
	} else if marks >= 70 {
		return "C"
	} else if marks >= 60 {
		return "D"
	} else {
		return "F"
	}
}
