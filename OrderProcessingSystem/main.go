package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var wg sync.WaitGroup
var lock sync.RWMutex
var errs []error
var orderCounter int

type Order struct {
	orderId int
	Items   []string
	Amount  float64
}

func (o *Order) String() string {
	return fmt.Sprintf("Order ID : %d\nOrder Items : %v\nOrder Amount : %.2f", o.orderId, o.Items, o.Amount)
}

type StockItem struct {
	name     string
	price    float64
	quantity int
}

func ProcessOrder(workerId int, orderChannel chan *Order, stockMap map[string]*StockItem) {
	defer wg.Done()
	for order := range orderChannel {
		amount, err := Checkout(order.Items, stockMap)
		order.Amount = amount
		if err != nil {
			errs = append(errs, err)
			fmt.Printf("[Worker-%d] Processing Order %d ....  ❌ Out of Stock\n", workerId, order.orderId)
		} else {
			var paymentStatus string
			_, err := getPayment(order.Amount)
			if err != nil {
				errs = append(errs, err)
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
				fmt.Printf("[Worker-%d] Processing Order %d ....  %s\n", workerId, order.orderId, paymentStatus)
				lock.Unlock()
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

func CreateOrder(orderID int, items []string, stockMap map[string]*StockItem) *Order {
	amount, _ := Checkout(items, stockMap)
	return &Order{
		orderId: orderID,
		Items:   items,
		Amount:  amount,
	}
}

func getPayment(amount float64) (bool, error) {
	random := rand.Int63n(3)
	if random == 0 {
		return false, fmt.Errorf("payment of rs %.2f failed", amount)
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

	order1 := CreateOrder(1, []string{"phone", "charger", "mouse"}, stockMap)
	order2 := CreateOrder(2, []string{"laptop"}, stockMap)
	order3 := CreateOrder(3, []string{"phone", "charger"}, stockMap)
	order4 := CreateOrder(4, []string{"mouse", "laptop", "mouse"}, stockMap)

	orders := []*Order{
		order3, order2, order1, order4,
	}

	fmt.Println("-------Orders-------")
	for _, order := range orders {
		fmt.Println(order)
		fmt.Println("-------------------")

	}
	fmt.Println()
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
	//	order := CreateOrder(orderCounter, items, stockMap)
	//	orderCounter++
	//	orderChannel <- order
	//}

	close(orderChannel)
	wg.Wait()

	fmt.Println()
	fmt.Println("-------Errors-------")
	for _, err := range errs {
		fmt.Printf("Error : %v \n", err)
	}
}
