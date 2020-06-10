/*
Problem: To represent a set of not more than 100 integers
as a process, S, which accepts two kinds of instruction
from its calling process X: (1) S!insert(n), insert the
integer n in the set, and (2) S!has(n); ... ; S?b, b is set true
if n is in the set, and false otherwise. The initial value of
the set is empty.
*/
package main

import (
	"fmt"
	"time"
)

type messageType uint8

const (
	insert messageType = iota
	has    messageType = iota
)

type operation struct {
	kind  messageType
	value uint8
}

func main() {
	messages := make(chan operation)
	hasElem := make(chan bool)
	go set(messages, hasElem)
	var i uint8
	for i = 0; i < 101; i++ {
		messages <- operation{insert, i}
	}
	// These must be goroutines, otherwise the `set` goroutine will be blocked trying to
	// send to hasElem since we don't recieve from it, causing a deadlock.
	go func() { messages <- operation{has, 3} }()
	go func() { messages <- operation{has, 5} }()
	go func() { messages <- operation{has, 100} }()
	fmt.Println(<-hasElem)
	fmt.Println(<-hasElem)
	fmt.Println(<-hasElem)
	time.Sleep(100 * time.Millisecond)
}

func set(messages <-chan operation, hasElem chan<- bool) {
	maxSize := uint8(100)
	data := make(map[uint8]bool)
	for message := range messages {
		switch message.kind {
		case insert:
			if message.value < maxSize {
				data[message.value] = true
			}
		case has:
			_, containsElem := data[message.value]
			hasElem <- containsElem
		}
	}
}
