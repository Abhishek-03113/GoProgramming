package main

import (
	conv "HandsOn2/Conversions"
	"fmt"
)

func main() {

	// Testing values

	weight := 70.00
	var length float64 = 10.00
	temperature := 25.00

	// type safe conversions using float64
	lengthInFeet := conv.Length(length)
	temperatureInFarenhite := conv.Temperature(temperature)
	weightInPounds := conv.Weight(weight)

	fmt.Printf("%.2f Celcius is %.2f Farenhite \n", temperature, temperatureInFarenhite)
	fmt.Printf("%.2f Meters is %.2f Feet\n", length, lengthInFeet)
	fmt.Printf("%.2f Kilograms is %.2f Pounds\n", weight, weightInPounds)

	var intLengthTest int = 10
	var intWeightTest int = 70
	intTemperatureTest := 25

	// Explicit type conversion
	lengthTest := conv.Length(float64(intLengthTest))
	weightTest := conv.Weight(float64(intWeightTest))
	temperatureTest := conv.Temperature(float64(intTemperatureTest))

	fmt.Printf("%d Celcius is %.2f Farenhite \n", intTemperatureTest, temperatureTest)
	fmt.Printf("%d Meters is %.2f Feet\n", intLengthTest, lengthTest)
	fmt.Printf("%d Kilograms is %.2f Pounds\n", intWeightTest, weightTest)

	num1 := 5
	num2 := 2.50

	fmt.Printf("Multiplying %d by %.2f is %.2f", num1, num2, (float64(num1))*num2)

}
