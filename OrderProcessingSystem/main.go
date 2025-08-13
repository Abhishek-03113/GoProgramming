package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"sync"
)

var wg sync.WaitGroup
var lock sync.RWMutex

type Order struct {
	orderId int
	Items   []string
	Amount  float64
}

func ProcessOrder(workerId int, orderChannel chan *Order, stockMap map[string][2]float64) {
	defer wg.Done()
	for order := range orderChannel {
		amount, _ := Checkout(order.Items, stockMap)
		order.Amount = amount
		if order.Amount == -1 {
			fmt.Printf("[Worker-%d] Processing Order %d ....  Out of Stock \n", workerId, order.orderId)

		} else {
			var paymentStatus string
			_, err := getPayment()
			if err != nil {
				paymentStatus = "Payment Failed"
				fmt.Printf("[Worker-%d] Processing Order %d ....  %s \n", workerId, order.orderId, paymentStatus)
			} else {
				paymentStatus = "✅ Payment Success"
				lock.Lock()
				for _, item := range order.Items {

					arr := stockMap[item]
					arr[0]--
					stockMap[item] = arr
				}
				lock.Unlock()
				fmt.Printf("[Worker-%d] Processing Order %d ....  %s \n", workerId, order.orderId, paymentStatus)

			}
		}
	}

	//fmt.Printf("Closing the order channel\n")
}

func Checkout(items []string, stockMap map[string][2]float64) (float64, error) {
	lock.RLock()
	defer lock.RUnlock()
	var bill float64
	for _, item := range items {
		if stockMap[item][0] <= 0 {
			return -1, fmt.Errorf("item %s is out of stock", item)
		}
		bill += stockMap[item][1]
	}

	return bill, nil
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
	return true, nil
}

func main() {
	stockMap := map[string][2]float64{
		"Charger": {10.00, 25.00},
		"Mouse":   {5.00, 10.00},
		"Laptop":  {3.00, 100.00},
		"Phone":   {1.00, 60.00},
	}

	orderChannel := make(chan *Order)

	for i := range 3 {
		wg.Add(1)
		go ProcessOrder(i, orderChannel, stockMap)
	}

	order1 := CreateOrder(1, []string{"Phone", "Charger", "Mouse"}, stockMap)
	order2 := CreateOrder(2, []string{"Laptop"}, stockMap)
	order3 := CreateOrder(3, []string{"Phone", "Charger"}, stockMap)
	order4 := CreateOrder(4, []string{"Mouse", "Laptop", "Mouse"}, stockMap)

	orders := []Order{
		*order3, *order2, *order1, *order4,
	}

	for i := 0; i < 4; i++ {
		orderChannel <- &orders[i]
	}

	close(orderChannel)
	wg.Wait()

}
