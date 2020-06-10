/*
Problem: Compute a factorial by the recursive method, to a given limit.

Notes: in the recursive formula fac(n) = n * fac(n-1). To accomplish this using
goroutines we can write a goroutine to compute the factorial that sends results back
over a channel. When we need to recurse we can just create new channels to recieve
intermediate results.
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	result := make(chan int)
	go factorial(result, 4)
	fac := <-result
	fmt.Println(fac)
	time.Sleep(100 * time.Millisecond)
}

func factorial(result chan int, n int) {
	switch n {
	case 0, 1:
		result <- 1
	default:
		subResult := make(chan int)
		go factorial(subResult, n-1)
		subFac := <-subResult
		result <- n * subFac
	}
}
