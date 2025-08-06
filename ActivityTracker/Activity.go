package main

import (
	"fmt"
	"time"
)

type Status int

const (
	NotStarted Status = iota
	Started
	Finished
)

type Activity struct {
	action string
	start  time.Time
	Time   time.Duration
	status Status
}

func NewActivity(action string) *Activity {
	return &Activity{action: action}
}

func (a *Activity) calcTime() time.Duration {
	return time.Since(a.start)
}

func (a *Activity) startActivity() {
	fmt.Printf("Activity %s started \n", a.action)
	a.status = Started
	a.start = time.Now()
}

func (a *Activity) stopActivity() {
	a.Time = a.calcTime()
	a.status = Finished
	fmt.Printf("Activity %s stopped, runtime\n", a.action)
}

func (a *Activity) String() string {
	return fmt.Sprintf("Activity %s started at %v and lasted for %v", a.action, a.start, a.Time)

}

func (a *Activity) Log() string {
	if a.status != Finished {
		fmt.Println("Activity Still running or Activity Not Started Yet")
		return ""
	}
	fmt.Println(a.String())
	return a.String()
}
