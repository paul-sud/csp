/*
Problem: Adapt the previous program to replace every
pair of consecutive asterisks "**" by an upward arrow
"↑". Assume that the final character input is not an
asterisk.

(2) As an exercise, adapt
this process to deal sensibly with input which ends with
an odd number of asterisks.

In Go, one solution is for west to close the channel to indicate to the copy goroutine
that there are no more values to be sent. Then the goroutine doesn't get blocked trying
to recieve the next value from the channel when west does not have any more characters
to send.
*/

package main

import (
	"fmt"
	"time"
)

func west(chars chan<- rune) {
	for _, char := range "he*l**lo***" {
		chars <- char
	}
	close(chars)
}

func east(chars <-chan rune) {
	for char := range chars {
		fmt.Println(string(char))
	}
}

func copy(input <-chan rune, output chan<- rune) {
	for char := range input {
		if char == '*' {
			next, ok := <-input
			if !ok {
				output <- char
			} else if next == '*' {
				output <- '↑'
			} else {
				output <- char
				output <- next
			}
		} else {
			output <- char
		}
	}
}

func main() {
	data := make(chan rune)
	output := make(chan rune)
	go west(data)
	go east(output)
	go copy(data, output)
	time.Sleep(100 * time.Millisecond)
}
