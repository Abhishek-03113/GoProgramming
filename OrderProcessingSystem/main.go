package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"sync"
)

var wg sync.WaitGroup
var lock sync.RWMutex

type Order struct {
	orderId int
	Items   []string // List of items in the order
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
				for _, itemName := range order.Items {
					if stockItem, exists := stockMap[itemName]; exists {
						stockItem.quantity--
					}
				}
				lock.Unlock()
				fmt.Printf("[Worker-%d] Processing Order %d ....  %s\n", workerId, order.orderId, paymentStatus)
			}
		}
	}
}

func Checkout(items []string, stockMap map[string]*StockItem) (float64, error) {
	lock.RLock()
	defer lock.RUnlock()
	var bill float64
	for _, itemName := range items {
		stockItem, exists := stockMap[itemName]
		if !exists || stockItem.quantity <= 0 {
			return -1, fmt.Errorf("item %s is out of stock", itemName)
		}
		bill += stockItem.price
	}
	return bill, nil
}

func CreateOrder(orderID int, items []string, stockMap map[string]*StockItem) *Order {
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
		"Phone":   {name: "Phone", price: 60.00, quantity: 2},
	}

	orderChannel := make(chan *Order)

	// Start worker goroutines
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go ProcessOrder(i, orderChannel, stockMap)
	}

	scanner := bufio.NewScanner(os.Stdin)
	orderCounter := 1

	for {
		fmt.Println("Create a new order (yes/no)?")
		scanner.Scan()
		newOrder := strings.ToLower(strings.TrimSpace(scanner.Text()))

		if newOrder != "yes" {
			close(orderChannel)
			break
		}

		// Collect items for the order
		var items []string
		for {
			fmt.Println("Enter item name (or 'done' to finish):")
			scanner.Scan()
			item := strings.TrimSpace(scanner.Text())
			if strings.ToLower(item) == "done" {
				break
			}
			items = append(items, item)
		}

		order := CreateOrder(orderCounter, items, stockMap)
		orderCounter++
		orderChannel <- order
	}

	wg.Wait()
	fmt.Println("All orders processed for today")
}
