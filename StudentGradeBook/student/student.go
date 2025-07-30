package student

import "fmt"

type Student struct {
	Name  string
	Age   int
	Grade string
}

func (s Student) PrintDetails() {
	fmt.Printf("Name: %s\nAge: %d\nGrade: %s\n", s.Name, s.Age, s.Grade)
}

func (s Student) UpdateGrade(newGrade string) {
	s.Grade = newGrade
}

func createStudent(name string, age int, grade string) Student {
	return Student{
		Name:  name,
		Age:   age,
		Grade: grade,
	}
}
