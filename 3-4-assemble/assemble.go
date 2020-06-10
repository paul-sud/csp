/*
Problem: To read a stream of characters from process X
and print them in lines of 125 characters on a lineprinter.
The last line should be completed with spaces if necessary. 
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

func processCard(cards <-chan string, output chan<- string) {
	for card := range cards {
		for _, char := range card {
			output <- char
		}
		output <- ' '
	}
}

func printStream(output <-chan string) {
	for line := range output {
		fmt.Println(line)
	}
}

func main() {
	cards := make(chan string)
	output := make(chan string)
	go drawCards(cards)
	go processCard(cards, output)
	go printStream(output)
	time.Sleep(100 * time.Millisecond)
}
