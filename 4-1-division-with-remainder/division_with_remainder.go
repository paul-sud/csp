/*
Problem: Construct a process to represent a function-
like subroutine, which accepts a positive dividend and
divisor, and returns their integer quotient and remainder.
Efficiency is of no concern.
*/
package main

import (
	"fmt"
	"time"
)

type result struct {
	quotient  int
	remainder int
}

func main() {
	results := make(chan result)
	go divideSub(results, 13, 3)
	res := <-results
	fmt.Println(res.quotient, res.remainder)
	time.Sleep(100 * time.Millisecond)
}

func divideSub(results chan<- result, dividend int, divisor int) {
	quotient := dividend / divisor
	remainder := dividend % divisor
	results <- result{quotient, remainder}
}
