/*
Problem: Adapt the previous program to replace every
pair of consecutive asterisks "**" by an upward arrow
"↑". Assume that the final character input is not an
asterisk.
*/

package main

import (
	"fmt"
	"time"
)

func west(chars chan<- rune) {
	for _, char := range "he*l**lo**g" {
		chars <- char
	}
}

func east(chars <-chan rune) {
	for char := range chars {
		fmt.Println(string(char))
	}
}

func copy(input <-chan rune, output chan<- rune) {
	for char := range input {
		if char == '*' {
			next := <-input
			if next == '*' {
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
