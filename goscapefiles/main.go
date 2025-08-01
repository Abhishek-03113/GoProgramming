package main

import "fmt"

func main() {

	m := map[string]int{"first": 1, "second": 2, "third": 3}
	i := 0
	for i < 3 {
		for k, v := range m {
			fmt.Printf("%s: %d \n", k, v)
		}
		i++
	}
}

//
//	temperature := 30
//`
//	kvmap := make(map[string]int)
//
//	kvmap["hello"] = 1
//
//	if temperature > 30 {
//		fmt.Print("Its hot outside")
//	} else if temperature < 20 {
//		fmt.Print("It's cold outside ")
//	} else {
//		fmt.Print("Its perfect weather today !!")
//	}
//
//	if userId := getId(); userId != -1 {
//		fmt.Print("Its Working")
//	} else {
//		fmt.Print("Not Working ")
//	}
//
//	describePlanet("EartEh")
//
//	switch day := 2; day {
//	case 1:
//		fmt.Println("Monday")
//		fallthrough
//	case 2:
//		fmt.Println("Tuesday")
//		fallthrough
//	case 3:
//		fmt.Println("Wednesday")
//	default:
//		fmt.Println("Another day")
//	}
//
//	processGrade("A+")
//
//	fmt.Println("Fibonacci of 5 is:", fibonacci(7))
//
//	cache := make(map[int]int)
//	fmt.Println("Factorial of 5 is:", factorial(5, cache))
//	fmt.Println("Factorial of 6 is:", factorial(6, cache))
//	fmt.Println("Factorial of 7 is:", factorial(7, cache))
//	fmt.Println("Factorial of 8 is:", factorial(8, cache))
//	fmt.Println("Factorial of 9 is:", factorial(9, cache))
//	fmt.Println("Factorial of 10 is:", factorial(10, cache))
//}
//
//func getId() int {
//	return 100
//}
//
//func describePlanet(planet string) {
//	switch planet {
//	case "Mercury":
//		fmt.Print("first one")
//	case "Venus":
//		fmt.Print("Second one")
//	case "Earth":
//		fmt.Print("Third Planet")
//	default:
//		fmt.Print("this is not a planet")
//	}
//}
//
//func processGrade(grade string) {
//	switch grade {
//	case "A+":
//		fmt.Println("Perfect score!")
//		fallthrough // Continues to next case
//	case "A":
//		fmt.Println("Excellent work!")
//		fmt.Println("You're on the honor roll!")
//	case "B":
//		fmt.Println("Good job!")
//	case "C":
//		fmt.Println("Satisfactory")
//	default:
//		fmt.Println("Needs improvement")
//	}
//}
//
//func fibonacci(n int) int {
//	if n <= 0 {
//		return 0
//	} else if n == 1 {
//		return 1
//	} else {
//		return fibonacci(n-1) + fibonacci(n-2)
//	}
//}
//
//func factorial(n int, cache map[int]int) int {
//	if val, found := cache[n]; found {
//		return val
//	}
//	if n <= 1 {
//		return 1
//	} else {
//		result := n * factorial(n-1, cache)
//		cache[n] = result
//		return result
//	}
//}
