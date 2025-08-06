package main

func main() {
	//
	aliceAccount := Account{
		HolderName: "Alice",
		Balance:    100,
	}

	// Deposite
	aliceAccount.Display()

	aliceAccount.Deposit(200)

	aliceAccount.Display()

	//withdraw

	aliceAccount.Withdraw(500)
	aliceAccount.Withdraw(25)

	aliceDebitCard := aliceAccount.IssueCard("1245")

	aliceDebitCard.debit(250)
	//r1 := Rectangle{
	//	Lenth: 20.00,
	//	Width: 20.00,
	//}
	//
	//r2 := Rectangle{
	//	Lenth: 12,
	//	Width: 6,
	//}
	//
	//fmt.Printf("Area of rectangle r1 %.2f and Perimiter of rectangle is %2.f \n", r1.area(), r1.perimeter())
	//fmt.Printf("Area of rectangle r2 %.2f and Perimiter of rectangle is %2.f \n", r2.area(), r2.perimeter())

	//abhishekEmail := EmailUser{
	//	email: "abhishek@gmail.com",
	//}
	//
	//abhishekSMS := SMSUser{
	//	phone: "9420668042",
	//}
	//
	//notifierSlice := []Notifier{
	//	abhishekEmail, abhishekSMS,
	//}
	//
	//for _, notifier := range notifierSlice {
	//	fmt.Println(notifier.Notify())
	//}

}

type Rectangle struct {
	Lenth float64
	Width float64
}

func (r Rectangle) area() float64 {
	return r.Lenth * r.Width
}

func (r Rectangle) perimeter() float64 {

	return 2 * (r.Lenth + r.Width)

}
