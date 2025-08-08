package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Central Monitor")

	tempChan := make(chan float64, 10)
	windChan := make(chan float64, 10)
	humiChan := make(chan float64, 10)

	go TemperatureSensor(tempChan)
	go WindSensor(windChan)
	go HumiditySensor(humiChan)

	for tempChan != nil && windChan != nil && humiChan != nil {
		select {
		case tempReading, ok := <-tempChan:
			if ok {
				fmt.Println("Reading from temperature sensor : ", tempReading)
			} else {
				fmt.Println("temp channel closed")
				tempChan = nil // Prevent further reads from tempChan
			}
		case windReading, ok := <-windChan:
			if ok {
				fmt.Println("Reading from wind sensor : ", windReading)
			} else {
				fmt.Println("Wind Channel is closed")
				windChan = nil // Prevent further reads from tempChan
			}
		case humidityReading, ok := <-humiChan:
			if ok {
				fmt.Println("Reading from Humidity sensor : ", humidityReading)
			} else {
				fmt.Println("Humidity Channel Closed")
				humiChan = nil // Prevent further reads from windChan
			}
		case <-time.After(4 * time.Second):
			fmt.Println("No readings received in the last second")

		}

	}

}

func TemperatureSensor(tempChan chan<- float64) {
	temperatures := []float64{32.5, 30.5, 40.9, 42.6}
	n := len(temperatures)
	for idx, temp := range temperatures {
		tempChan <- temp
		fmt.Println("Remaining temperature readings: ", n-idx-1)
		time.Sleep(2 * time.Second)
	}

	close(tempChan)
}

func HumiditySensor(humiditychan chan<- float64) {
	humidities := []float64{12.5, 13.5, 15.9, 19.6}
	n := len(humidities)
	for idx, humi := range humidities {
		humiditychan <- humi
		fmt.Println("Remaining humidity readings: ", n-idx-1)
		time.Sleep(time.Second)
	}
	close(humiditychan)

}

func WindSensor(windchan chan<- float64) {
	windspeeds := []float64{112.5, 113.5, 115.9, 119.6}
	n := len(windspeeds)
	for idx, wind := range windspeeds {
		windchan <- wind
		fmt.Println("Remaining wind readings: ", n-idx-1)
		time.Sleep(3 * time.Second)
	}
	close(windchan)

}
