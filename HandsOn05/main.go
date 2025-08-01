package main

import "fmt"
import "HandsOn05/ScoreRating"

func main() {

	//row, col := 10, 10
	//
	//for i := 1; i <= row; i++ {
	//	for j := 1; j <= col; j++ {
	//		fmt.Printf("%d x %d = %d \n", i, j, i*j)
	//	}
	//}

	testScores := []float64{76, 43, 49, 99, 86, 78, 90, 96, 42, 12, 66, 67}
	//testScoresInt := []int{76, 43, 49, 99, 86, 78, 90, 96, 42, 12, 66, 67}
	//
	//for _, score := range testScoresInt {
	//	fmt.Printf("Score : %d, Remark : %s \n", score, ScoreRating.ScoreRating(score))
	//}

	fmt.Printf("Empty args test %.4f\n", ScoreRating.Average())
	fmt.Printf("Average of 4,5,6,10 is %.4f\n", ScoreRating.Average(4, 5, 6, 10))
	fmt.Printf("Average of TestScores is %.4f", ScoreRating.Average(testScores...))
}
