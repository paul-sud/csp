/*
Problem: Write a process X to copy characters output by
process west to process, east.
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	data := make(chan rune)
	output := make(chan rune)
	go west(data)
	go east(output)
	go copy(data, output)
	time.Sleep(100 * time.Millisecond)
}

func west(chars chan<- rune) {
	for _, char := range "hello" {
		chars <- char
	}
}

func east(chars <-chan rune) {
	for char := range chars {
		fmt.Println(string(char))
	}
}

func copy(input <-chan rune, output chan<- rune) {
	for buf := range input {
		output <- buf
	}
}
