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
	for _, char := range "he*l**lo***" {
		chars <- char
	}
}

func east(chars <-chan rune) {
	for char := range chars {
		fmt.Println(string(char))
	}
}

func copy(input <-chan rune, output chan<- rune) {
	prevWasAsterisk := false
	for char := range input {
		if char == '*' {
			if prevWasAsterisk {
				output <- '↑'
				prevWasAsterisk = false
			} else {
				prevWasAsterisk = true
				continue
			}
		} else {
			if prevWasAsterisk {
				output <- '*'
			}
			output <- char
			prevWasAsterisk = false
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
	fmt.Printf("\n")
}
