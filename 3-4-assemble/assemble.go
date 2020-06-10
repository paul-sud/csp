/*
Problem: To read a stream of characters from process X
and print them in lines of 125 characters on a lineprinter.
The last line should be completed with spaces if necessary.
*/
package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func generateRandomChars(chars chan<- string) {
	rand.Seed(42)
	for i := 0; i < 255; i++ {
		randString := strconv.Itoa(rand.Intn(127))
		chars <- randString
	}
	close(chars)
}

func processChars(chars <-chan string, output chan<- []string) {
	i := 0
	var b [125]string
	for j := range b {
		b[j] = " "
	}
	for char := range chars {
		fmt.Println(char)
		b[i] = char
		if i == 124 {
			output <- b[:]
			i = 0
			for j := range b {
				b[j] = " "
			}
		}
		i++
	}
	output <- b[:]
}

func linePrinter(output <-chan []string) {
	// for line := range output {
	// 	fmt.Printf(string(line))
	// }
}

func main() {
	chars := make(chan string)
	output := make(chan []string)
	go generateRandomChars(chars)
	go processChars(chars, output)
	go linePrinter(output)
	time.Sleep(100 * time.Millisecond)
}
