package main

import (
	"fmt"
	"time"
)

type Activity struct {
	action string
	start  time.Time
	Time   time.Duration
	status bool
}

func NewActivity(action string) *Activity {
	return &Activity{action: action}
}

func (a *Activity) calcTime() time.Duration {
	return time.Since(a.start)
}

func (a *Activity) startActivity() {
	fmt.Printf("Activity %s started", a.action)
	a.status = true
	a.start = time.Now()
}

func (a *Activity) stopActivity() {
	fmt.Printf("Activity %s stopped", a.action)
	a.calcTime()
	a.status = false
}

func (a *Activity) String() string {
	return fmt.Sprintf("Activity %s started at %v and lasted for %v", a.action, a.start, a.Time)

}

func (a *Activity) Log() string {
	if a.status {
		return "Activity Still running"
	}
	fmt.Println(a.String())
	return a.String()

}
