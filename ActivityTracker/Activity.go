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
	Action string
	Start  time.Time
	Time   time.Duration
	Status Status
}

func NewActivity(Action string) *Activity {
	return &Activity{Action: Action, Status: NotStarted}
}

func (a *Activity) calcTime() time.Duration {
	return time.Since(a.Start)
}

func (a *Activity) startActivity() {
	a.Start = time.Now()
	a.Status = Started
}

func (a *Activity) stopActivity() {
	a.Time = a.calcTime()
	a.Status = Finished
}

func (a *Activity) String() string {

	if a.Status == Finished {
		return fmt.Sprintf("Activity %s started at %v:%v:%v and lasted for %v \n", a.Action, a.Start.Hour(), a.Start.Minute(), a.Start.Second(), a.Time)
	}
	return fmt.Sprintf("Activity %s started at %v:%v:%v and running since %v \n", a.Action, a.Start.Hour(), a.Start.Minute(), a.Start.Second(), a.calcTime())
}

func (a *Activity) Log() string {
	return a.String()
}
