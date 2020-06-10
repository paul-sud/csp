/*
Problem: to read cards from a cardfile and output to
process X the stream of characters they contain. An extra
space should be inserted at the end of each card.
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	cards := make(chan string)
	output := make(chan rune)
	go drawCards(cards)
	go processCard(cards, output)
	go printStream(output)
	time.Sleep(100 * time.Millisecond)
}

func drawCards(cards chan<- string) {
	rand.Seed(42)
	for i := 0; i < 2; i++ {
		card := make([]rune, 80)
		for i := range card {
			card[i] = rune(33 + rand.Intn(93))
		}
		cards <- string(card)
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
