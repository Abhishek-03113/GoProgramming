package main

import (
	"errors"
	"fmt"
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
		stockItem, err := stockMap[itemName]
		if !err || stockItem.quantity <= 0 {
			return -1, fmt.Errorf("item %s is out of stock", itemName)
		}
		bill += stockItem.price
	}
	return bill, nil
}

func CreateOrder(orderID int, items []string) *Order {
	return &Order{
		orderId: orderID,
		Items:   items,
	}
}

func getPayment() (bool, error) {
	random := rand.Int63n(4)
	if random == 0 {
		return false, errors.New("payment failed")
	}
	return true, nil
}

func main() {

	stockMap := map[string]*StockItem{
		"charger": {name: "charger", price: 25.00, quantity: 10},
		"mouse":   {name: "mouse", price: 10.00, quantity: 5},
		"laptop":  {name: "laptop", price: 100.00, quantity: 2},
		"phone":   {name: "phone", price: 60.00, quantity: 1},
	}

	orderChannel := make(chan *Order)

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go ProcessOrder(i, orderChannel, stockMap)
	}

	order1 := CreateOrder(1, []string{"phone", "charger", "mouse"})
	order2 := CreateOrder(2, []string{"laptop"})
	order3 := CreateOrder(3, []string{"phone", "charger"})
	order4 := CreateOrder(4, []string{"mouse", "laptop", "mouse"})

	orders := []*Order{
		order3, order2, order1, order4,
	}

	for _, order := range orders {
		orderChannel <- order
	}
	//scanner := bufio.NewScanner(os.Stdin)
	//orderCounter := 1
	//
	//for {
	//	fmt.Printf("Create a new order (yes/no)?")
	//	scanner.Scan()
	//	newOrder := strings.ToLower(strings.TrimSpace(scanner.Text()))
	//
	//	if newOrder != "yes" {
	//		close(orderChannel)
	//		break
	//	}
	//
	//	var items []string
	//	var itemscsv string
	//
	//	fmt.Println("input the items you want to buy in a single line separated by comma : ")
	//	scanner.Scan()
	//
	//	itemscsv = strings.ToLower(strings.TrimSpace(scanner.Text()))
	//	itemslist := strings.Split(itemscsv, ",")
	//
	//	for _, item := range itemslist {
	//		items = append(items, strings.TrimSpace(item))
	//	}
	//	order := CreateOrder(orderCounter, items)
	//	orderCounter++
	//	orderChannel <- order
	//}
	close(orderChannel)
	wg.Wait()
	fmt.Println("All orders processed for today")
}
