package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
)

type Order struct {
	orderId int
	Items   []string
	Amount  float64
}

func ProcessOrder(workerId int, orderChannel chan Order) {
	for order := range orderChannel {
		if order.Amount == -1 {
			fmt.Printf("[Worker-%d] Processing Order %d ....  Out of Stock", workerId, order.orderId)
			continue
		}
		var paymentStatus string
		_, err := getPayment()
		if err != nil {
			paymentStatus = "Payment Failed"
			log.Println(err)
		} else {
			paymentStatus = "Payment Successful"
		}
		fmt.Printf("[Worker-%d] Processing Order %d ....  %s", workerId, order.orderId, paymentStatus)
	}
}

func Checkout(items []string, stockMap map[string][2]float64) (float64, error) {
	var bill float64
	for _, item := range items {
		if stockMap[item][0] < 0 {
			return -1, fmt.Errorf("item %s is out of stock", item)
		}
		bill += stockMap[item][1]
	}

	return bill, _
}

func CreateOrder(orderID int, Items []string, stockMap map[string][2]float64) *Order {
	amount, err := Checkout(Items, stockMap)
	if err != nil {
		log.Println(err)
	}

	return &Order{
		orderId: orderID,
		Items:   Items,
		Amount:  amount,
	}
}

func getPayment() (bool, error) {
	random := rand.Int63n(10)
	if random == 0 {
		return false, errors.New("payment failed")
	}
	return true, _
}

func main() {
	stockMap := map[string][2]float64{
		"Charger": {10.00, 25.00},
		"Mouse":   {5.00, 10.00},
		"Laptop":  {3.00, 100.00},
		"Phone":   {1.00, 60.00},
	}

	orderChannel := make(chan Order)

}
