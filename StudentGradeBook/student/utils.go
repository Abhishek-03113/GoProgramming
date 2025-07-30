package student

func CalculateTopStudent(students []Student) Student {
	if len(students) == 0 {
		return Student{}
	}

	topStudent := students[0]
	for _, student := range students {
		if student.Marks > topStudent.Marks {
			topStudent = student
		}
	}
	return topStudent
}

func CalculateAverageMarks(students []Student) int {
	averageMarks := 0

	for _, student := range students {
		averageMarks += student.Marks
	}

	return averageMarks / len(students)
}

func CalculateBestMarks(students []Student) int {

	return CalculateTopStudent(students).Marks

}
