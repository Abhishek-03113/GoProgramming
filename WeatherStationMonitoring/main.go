package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Central Monitor")

	tempChan := make(chan float64)
	windChan := make(chan float64)
	humiChan := make(chan float64)

	go TemperatureSensor(tempChan)
	go WindSensor(windChan)
	go HumiditySensor(humiChan)

	tcnt, wcnt, hcnt := 0, 0, 0
	for {
		select {
		case tempReading, ok := <-tempChan:
			if ok {
				fmt.Printf("Temperature sensor, cnt : %d, reading : %.2f \n", tcnt, tempReading)
				tcnt++
			} else {
				fmt.Println("temp channel closed")
				tempChan = nil
			}
		case windReading, ok := <-windChan:
			if ok {
				fmt.Printf("Wind sensor, cnt : %d, reading : %.2f \n", wcnt, windReading)
				wcnt++
			} else {
				fmt.Println("Wind Channel is closed")
				windChan = nil
			}
		case humidityReading, ok := <-humiChan:
			if ok {
				fmt.Printf("Humidity sensor, cnt : %d, reading : %.2f \n", hcnt, humidityReading)
				hcnt++
			} else {
				fmt.Println("Humidity Channel Closed")
				humiChan = nil
			}
		case <-time.After(4 * time.Second):
			fmt.Println("No readings received in the last second")
			return

		}

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
