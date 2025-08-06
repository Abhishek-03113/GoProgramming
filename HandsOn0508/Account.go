package main

import (
	"fmt"
	"strings"
)

type Validator interface {
	Validate(T any) bool
}

type SuffiecientBalanceValidator struct {
}

func (s SuffiecientBalanceValidator) Validate(input struct{ Balance, Amount float64 }) bool {
	return input.Balance-input.Amount > 0
}

type NameValidator struct{}

func (n NameValidator) Validate(input string) bool {
	return len(strings.Trim(input, " ")) > 0
}

type Account struct {
	HolderName string
	Balance    float64
}

func (a *Account) Deposit(amount float64) {
	a.Balance = a.Balance + amount
}

func (a *Account) Withdraw(amount float64) bool {
	validator := SuffiecientBalanceValidator{}
	if validator.Validate(struct{ Balance, Amount float64 }{Balance: a.Balance, Amount: amount}) {
		a.Balance -= amount
		fmt.Printf("Amount %.2f debitted, from %+v", amount, a)
		return true
	}
	fmt.Println("Insufficient Account Balance")
	return false
}

func (a *Account) Display() {
	fmt.Println(a)

}

func (a *Account) Validate(amount float64) bool {
	return a.Balance-amount > 0

}

func (a Account) IssueCard(number string) DebitCard {
	return DebitCard{
		Acc:    a,
		Number: number,
	}

}

func (a Account) String() string {
	return fmt.Sprintf("Account Holder Name %s, Account Balance %.2f \n", a.HolderName, a.Balance)
}
