/*
Problem: To print in ascending order all primes less than
10000. Use an array of processes, SIEVE, in which each
process inputs a prime from its predecessor and prints it.
The process then inputs an ascending stream of numbers
from its predecessor and passes them on to its successor,
suppressing any that are multiples of the original prime.

Notes: we can just feed the stream of numbers via sequential goroutines. Each goroutine
pops off the first value in the stream and prints it, then sends the remaining values
down the chain if they are not divisible by the first number it recieved. Note that once
the first value is larger than sqrt(N) where N is the largest value in the sequence, we
can just print the remaining values in the sequence without needing to spin up more
goroutines.
*/
package main

import (
	"fmt"
	"time"
)

// const numNums uint16 = 10000
// const sqrtNumNums uint16 = 100
const numNums uint16 = 100
const sqrtNumNums uint16 = 10

func main() {
	next := make(chan uint16)
	go sieve(next)
	var i uint16
	for i = 2; i < numNums; i++ {
		next <- i
	}
	time.Sleep(100 * time.Millisecond)
}

func sieve(nums <-chan uint16) {
	firstNum := <-nums
	fmt.Println(firstNum)
	if firstNum < sqrtNumNums {
		output := make(chan uint16)
		go sieve(output)
		for num := range nums {
			if num%firstNum != 0 {
				output <- num
			}
		}
	} else {
		for num := range nums {
			fmt.Println(num)
		}
	}
}
