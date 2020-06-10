/*
Problem: To read a stream of characters from process X
and print them in lines of 125 characters on a lineprinter.
The last line should be completed with spaces if necessary.
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	chars := make(chan rune)
	output := make(chan []rune)
	go generateRandomChars(chars)
	go processChars(chars, output)
	go linePrinter(output)
	time.Sleep(100 * time.Millisecond)
}

func generateRandomChars(chars chan<- rune) {
	rand.Seed(42)
	for i := 0; i < 256; i++ {
		randString := rune(33 + rand.Intn(93))
		chars <- randString
	}
	close(chars)
}

func processChars(chars <-chan rune, output chan<- []rune) {
	i := 0
	var b [125]rune
	for j := range b {
		b[j] = ' '
	}
	for char := range chars {
		b[i] = char
		if i == 124 {
			// Need to send a new copy of the slice every time, otherwise we will be
			// printing while the slice is being mutated, which leads to undefined
			// behavior
			tmp := make([]rune, 125)
			copy(tmp, b[:])
			output <- tmp
			for j := range b {
				b[j] = ' '
			}
			i = 0
		} else {
			i++
		}
	}
	output <- b[:]
}

func linePrinter(output <-chan []rune) {
	for line := range output {
		fmt.Println(string(line))
	}
}
