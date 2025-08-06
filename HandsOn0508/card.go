package main

type DebitCard struct {
	Acc    Account
	Number string
}

func (d DebitCard) debit(amount float64) bool {
	return d.Acc.Withdraw(amount)
}
