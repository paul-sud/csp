/*
Problem: Adapt the above program to replace every pair
of consecutive asterisks by an upward arrow.

Note that this is mostly just copy pasting functions from assemble, disassemble, and
squash.
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	cards := make(chan string)
	intermediate := make(chan rune)
	secondary := make(chan rune)
	output := make(chan []rune)
	go drawCards(cards)
	go addSpaceAfterCard(cards, intermediate)
	go squash(intermediate, secondary)
	go processChars(secondary, output)
	go linePrinter(output)
	time.Sleep(100 * time.Millisecond)
}

func drawCards(cards chan<- string) {
	rand.Seed(42)
	for i := 0; i < 4; i++ {
		card := make([]rune, 80)
		for i := range card {
			card[i] = rune(33 + rand.Intn(93))
		}
		asteriskIndex := rand.Intn(79)
		card[asteriskIndex] = '*'
		card[asteriskIndex+1] = '*'
		cards <- string(card)
	}
	close(cards)
}

func addSpaceAfterCard(cards <-chan string, output chan<- rune) {
	for card := range cards {
		for _, char := range card {
			output <- char
		}
		output <- ' '
	}
	close(output)
}

func squash(input <-chan rune, output chan<- rune) {
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
	close(output)
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
