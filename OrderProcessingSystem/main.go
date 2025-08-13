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
	Items   map[string]int // item name -> quantity
	Amount  float64
}

type StockItem struct {
	name     string
	price    float64
	quantity int
}

func ProcessOrder(workerId int, orderChannel chan *Order, stockMap map[string]*StockItem) {
	defer wg.Done()
	for order := range orderChannel {
		amount, _ := Checkout(order.Items, stockMap)
		order.Amount = amount
		if order.Amount == -1 {
			fmt.Printf("[Worker-%d] Processing Order %d ....  ❌ Out of Stock\n", workerId, order.orderId)
		} else {
			var paymentStatus string
			_, err := getPayment()
			if err != nil {
				paymentStatus = "❌ Payment Failed"
				fmt.Printf("[Worker-%d] Processing Order %d ....  %s\n", workerId, order.orderId, paymentStatus)
			} else {
				paymentStatus = "✅ Payment Success"
				lock.Lock()
				for itemName, qty := range order.Items {
					if stockItem, exists := stockMap[itemName]; exists {
						stockItem.quantity -= qty
					}
				}
				lock.Unlock()
				fmt.Printf("[Worker-%d] Processing Order %d ....  %s\n", workerId, order.orderId, paymentStatus)
			}
		}
	}
}

func Checkout(items map[string]int, stockMap map[string]*StockItem) (float64, error) {
	lock.RLock()
	defer lock.RUnlock()
	var bill float64
	for itemName, qty := range items {
		stockItem, exists := stockMap[itemName]
		if !exists || stockItem.quantity < qty {
			return -1, fmt.Errorf("item %s is out of stock", itemName)
		}
		bill += stockItem.price * float64(qty)
	}
	return bill, nil
}

func CreateOrder(orderID int, items map[string]int, stockMap map[string]*StockItem) *Order {
	amount, err := Checkout(items, stockMap)
	if err != nil {
		log.Println(err)
	}
	return &Order{
		orderId: orderID,
		Items:   items,
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
	stockMap := map[string]*StockItem{
		"Charger": {name: "Charger", price: 25.00, quantity: 10},
		"Mouse":   {name: "Mouse", price: 10.00, quantity: 5},
		"Laptop":  {name: "Laptop", price: 100.00, quantity: 3},
		"Phone":   {name: "Phone", price: 60.00, quantity: 1},
	}

	orderChannel := make(chan *Order)

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go ProcessOrder(i, orderChannel, stockMap)
	}

	order1 := CreateOrder(1, map[string]int{"Phone": 1, "Charger": 1, "Mouse": 1}, stockMap)
	order2 := CreateOrder(2, map[string]int{"Laptop": 1}, stockMap)
	order3 := CreateOrder(3, map[string]int{"Phone": 1, "Charger": 1}, stockMap)
	order4 := CreateOrder(4, map[string]int{"Mouse": 2, "Laptop": 1}, stockMap)

	orders := []*Order{order3, order2, order1, order4}

	for _, ord := range orders {
		orderChannel <- ord
	}

	close(orderChannel)
	wg.Wait()
}
