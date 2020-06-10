/*
Problem: to read cards from a cardfile and output to
process X the stream of characters they contain. An extra
space should be inserted at the end of each card.
*/
package main

import (
	"fmt"
	"time"
)

func drawCards(cards chan<- string) {
	data := []string{"foo", "bar"}
	for _, card := range data {
		cards <- card
	}
}

func processCard(cards <-chan string, output chan<- rune) {
	for card := range cards {
		for _, char := range card {
			output <- char
		}
		output <- ' '
	}
}

func printStream(output <-chan rune) {
	for char := range output {
		fmt.Println(string(char))
	}
}

func main() {
	cards := make(chan string)
	output := make(chan rune)
	go drawCards(cards)
	go processCard(cards, output)
	go printStream(output)
	time.Sleep(100 * time.Millisecond)
}
