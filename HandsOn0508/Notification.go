package main

import "fmt"

type Notifier interface {
	Notify() string
}

type EmailUser struct {
	email string
}

type SMSUser struct {
	phone string
}

func (s SMSUser) Notify() string {
	return fmt.Sprintf("Sent SMS notification User on the phone %s", s.phone)
}

func (e EmailUser) Notify() string {
	return fmt.Sprintf("Sent Email notification to Uer on email %s", e.email)
}
