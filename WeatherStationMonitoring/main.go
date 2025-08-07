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

	for i := time.Second; i < 30*time.Second; i += time.Second {
		select {
		case tempReading, ok := <-tempChan:
			if ok {
				fmt.Println("Reading from temperature sensor : ", tempReading)
			} else {
				fmt.Println("temp channel closed")
			}
		case windReading, ok := <-windChan:
			if ok {
				fmt.Println("Reading from wind sensor : ", windReading)
			} else {
				fmt.Println("Wind Channel is closed")
			}
		case humidityReading, ok := <-humiChan:
			if ok {
				fmt.Println("Reading from Humidity sensor : ", humidityReading)
			} else {
				fmt.Println("Humidity Channel Closed")
			}

		}

		fmt.Println(i)
	}

}

func TemperatureSensor(tempChan chan<- float64) {
	temperatures := []float64{32.5, 30.5, 40.9, 42.6}

	for _, temp := range temperatures {
		tempChan <- temp
		time.Sleep(2 * time.Second)
	}

	close(tempChan)
}

func HumiditySensor(humiditychan chan<- float64) {
	humidities := []float64{12.5, 13.5, 15.9, 19.6}
	for _, humi := range humidities {
		humiditychan <- humi
		time.Sleep(time.Second)
	}
	close(humiditychan)

}

func WindSensor(windchan chan<- float64) {
	windspeeds := []float64{112.5, 113.5, 115.9, 119.6}
	for _, wind := range windspeeds {
		windchan <- wind
		time.Sleep(3 * time.Second)
	}
	close(windchan)

}
