/*
Problem: To implement an integer semaphore, S, shared
among an array X(i:I..100) of client processes. Each
process may increment the semaphore by S!V() or
decrement it by S!P(), but the latter command must be
delayed if the value of the semaphore is not positive.

Note: if you have more decrements than increments, this will loop forever waiting for
increments! Try playing with the bool generator in line 57
*/
package main

import (
	"fmt"
	"sync"
	"math/rand"
)

func main() {
	rand.Seed(42)
	var wg sync.WaitGroup
	// We need to buffer the decrement channel to be able to handle loopback
	increment := make(chan bool)
	decrement := make(chan bool, 10)
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go signal(increment, decrement)
	}
	go semaphore(increment, decrement, &wg)
	wg.Wait()
}

func semaphore(increment <-chan bool, decrement chan bool, wg *sync.WaitGroup) {
	var value uint8
	value = 0
	for {
		select {
		case <-increment:
			value++
			fmt.Println("Recieved increment signal and successfully incremented")
			wg.Done()
		case <-decrement:
			if value > 0 {
				value--
				fmt.Println("Recieved decrement signal and successfully decremented")
				wg.Done()
			} else {
				decrement <- true
				fmt.Println("Recieved decrement signal but postponed decrement")
			}
		}
	}
}

func signal(increment chan<- bool, decrement chan<- bool) {
	// Try to increment most of the time so we are unlikely to trigger an infinite loop
	shouldIncrement := rand.Intn(10) > 3
	if shouldIncrement {
		fmt.Println("Sent increment signal")
		increment <- true
	} else {
		fmt.Println("Sent decrement signal")
		decrement <- true
	}
}
