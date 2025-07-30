package main

import (
	eval "HandsOn3/ExpressionEvaluator"
	"fmt"
)

func main() {

	var x, y float64

	outputMap := make(map[bool]string)
	outputMap[true] = "YES"
	outputMap[false] = "NO"

	fmt.Println("Input the value of X :")
	fmt.Scanf("%f", &x)
	fmt.Println("Input the value of Y :")
	fmt.Scanf("%f", &y)

	fmt.Println("Arithmetic Operations")

	fmt.Printf("Additon of %.2f and %.2f == %.2f \n ", x, y, eval.Add(x, y))
	fmt.Printf("Subtraction of %.2f and %.2f == %.2f \n ", x, y, eval.Minus(x, y))
	fmt.Printf("Product of %.2f and %.2f == %.2f \n ", x, y, eval.Multiply(x, y))
	fmt.Printf("Division of %.2f and %.2f == %.2f \n ", x, y, eval.Divide(x, y))
	fmt.Printf("Modulo of %.2f and %.2f == %d \n ", x, y, eval.Modulo(x, y))

	fmt.Println("---------------------------------")
	fmt.Println("Comparison Operators ")

	fmt.Printf("Equality of %.2f and %.2f == %v \n ", x, y, outputMap[eval.Equals(x, y)])
	fmt.Printf("Inequality of %.2f and %.2f == %v \n ", x, y, outputMap[eval.NotEquals(x, y)])
	fmt.Printf("is %.2f > %.2f == %v \n ", x, y, outputMap[eval.GreaterThan(x, y)])
	fmt.Printf("is %.2f < %.2f == %v \n ", x, y, outputMap[eval.LessThan(x, y)])
	fmt.Printf("is %.2f >= %.2f == %v \n ", x, y, outputMap[eval.GreaterThanEquals(x, y)])
	fmt.Printf("is %.2f <= %.2f == %v \n ", x, y, outputMap[eval.LessThanEquals(x, y)])

	fmt.Println("All operations completed successfully.")

}
