/*
5.3 Dining Philosophers (Problem due to E.W. Dijkstra)
Problem: Five philosophers spend their lives thinking
and eating. The philosophers share a common dining
room where there is a circular table surrounded by five
chairs, each belonging to one philosopher. In the center
of the table there is a large bowl of spaghetti, and the
table is laid with five forks (see Figure 1). On feeling
hungry, a philosopher enters the dining room, sits in his
own chair, and picks up the fork on the left of his place.
Unfortunately, the spaghetti is so tangled that he needs
to pick up and use the fork on his right as well. When he
has finished, he puts down both forks, and leaves the
room. The room should keep a count of the number of
philosophers.

Notes: (1) The solution given above does not prevent all
five philosophers from entering the room, each picking
up his left fork, and starving to death because he cannot
pick up his right fork. (2) Exercise: Adapt the above
program to avert this sad possibility. Hint: Prevent more
than four philosophers from entering the room. (Solution
due to E. W. Dijkstra)
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

const numPhilosophers int = 5
const maxThinkTime int = 1000
const eatTime int = 50
const simulationTimeSec int = 3

func main() {
	rand.Seed(42)
	roomEnter := make(chan bool, numPhilosophers)
	roomLeave := make(chan bool, numPhilosophers)
	forkPickup := make([]chan bool, 5)
	for i := 0; i < len(forkPickup); i++ {
		forkPickup[i] = make(chan bool)
	}
	forkPutdown := make([]chan bool, 5)
	for i := 0; i < len(forkPutdown); i++ {
		forkPutdown[i] = make(chan bool)
	}
	for i := 0; i < numPhilosophers; i++ {
		go philosopher(i, roomEnter, roomLeave, forkPickup, forkPutdown)
		go fork(forkPickup[i], forkPutdown[i])
	}
	go room(roomEnter, roomLeave)
	time.Sleep(time.Duration(simulationTimeSec) * time.Second)
}

func philosopher(chairNumber int, roomEnter chan<- bool, roomLeave chan<- bool, forkPickup []chan bool, forkPutdown []chan bool) {
	for {
		// Philosopher is thinking
		time.Sleep(time.Duration(rand.Intn(maxThinkTime)) * time.Millisecond)
		roomEnter <- true
		// Wait until both forks are acquired, then eat. Philosophers eat very quickly!
		// We can only pick up the forks on either side of the philospher's designated
		// chair, need to wait to acquire both forks. We don't want to try to just
		// sequentially receive from the left fork and then the right fork, because
		// the right fork might have been available in the meantime.
		select {
		case forkPickup[chairNumber] <- true:
			forkPickup[(chairNumber+1)%numPhilosophers] <- true
		case forkPickup[(chairNumber+1)%numPhilosophers] <- true:
			forkPickup[chairNumber] <- true
		}
		fmt.Println("I am eating:", chairNumber)
		time.Sleep(time.Duration(eatTime) * time.Millisecond)
		forkPutdown[chairNumber] <- true
		forkPutdown[chairNumber%numPhilosophers] <- true
		// Done eating
		roomLeave <- true
	}
}

func room(roomEnter <-chan bool, roomLeave <-chan bool) {
	count := 0
	for {
		switch {
		case count >= 4:
			select {
			case <-roomLeave:
				count--
				fmt.Println("Philosopher left the room, count is", count)
			}
		default:
			select {
			case <-roomEnter:
				count++
				fmt.Println("Philosopher entered the room, count is", count)
			case <-roomLeave:
				count--
				fmt.Println("Philosopher left the room, count is", count)
			}
		}
	}
}

func fork(forkPickUp <-chan bool, forkPutDown <-chan bool) {
	// Once picked up, blocks until put down, and vice versa
	for {
		<-forkPickUp
		<-forkPutDown
	}
}
