package main

import (
	"fmt"
	"sync"
)

var lock sync.Mutex

func bookTicket(id int, bookingChannel <-chan int, ticketChan chan<- string, wg *sync.WaitGroup, Tickets *int) {
	defer wg.Done()

	for requiredTicket := range bookingChannel {
		if requiredTicket == 0 {
			continue
			//ticketChan <- fmt.Sprintf("U cant buy 0 tickets \n")
		} else if *Tickets-requiredTicket > 0 {
			lock.Lock()
			*Tickets -= requiredTicket
			ticketChan <- fmt.Sprintf("%d booked %d tickets \n", id, requiredTicket)
			lock.Unlock()
		} else {
			ticketChan <- fmt.Sprintf("%d tickets are not available, available tickets %d \n", requiredTicket, *Tickets)
		}
	}

}

func main() {

	var wg sync.WaitGroup

	bookChan := make(chan int)
	ticketChan := make(chan string)

	totalTickets := 100
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go bookTicket(i, bookChan, ticketChan, &wg, &totalTickets)
	}

	for i := 1; i < 20; i++ {
		bookChan <- i
		fmt.Println(<-ticketChan)
	}

	close(bookChan)
	wg.Wait()
	close(ticketChan)

}
