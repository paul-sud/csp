/*
Problem: Extend the solution to 4.3 by providing a fast
method for scanning all members of the set without
changing the value of the set. The user program will
contain a repetitive command of the form:
S!scan( ); more:boolean; more := true;
* [more;x:integer; S?next(x) ->, ... deal with x ....
[]more; S?noneleft( )->, more := false
]
where S!scan( ) sets the representation into a scanning
mode. The repetitive command serves as a for statement,
inputting the successive members of x from the set and
inspecting them until finally the representation sends a
signal that there are no members left. The body of the
repetitive command is not permitted to communicate
with S in any way.

Note: I think this solution satisfies the last sentence of the problem statement.
Outside of the loop we trigger the scan operation, which causes S to send values over
a channel. Then all we do is recieve over the channel, without sending additional
messages to S.
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
	scan   messageType = iota
)

type operation struct {
	kind  messageType
	value uint8
}

type scanResult struct {
	value  uint8
	closed bool
}

func main() {
	messages := make(chan operation)
	hasElem := make(chan bool)
	scanResults := make(chan scanResult)
	go set(messages, hasElem, scanResults)
	var i uint8
	for i = 0; i < 10; i++ {
		messages <- operation{insert, i}
	}
	// These must be goroutines, otherwise the `set` goroutine will be blocked trying to
	// send to hasElem since we don't recieve from it, causing a deadlock.
	go func() { messages <- operation{scan, 0} }()
	// We don't want to close the channel, otherwise we can't scan multiple times.
	for val := range scanResults {
		if val.closed {
			break
		}
		fmt.Println(val.value)
	}
	time.Sleep(100 * time.Millisecond)
}

func set(messages <-chan operation, hasElem chan<- bool, scanValues chan<- scanResult) {
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
		case scan:
			for key := range data {
				scanValues <- scanResult{key, false}
			}
			scanValues <- scanResult{0, true}
		}
	}
}
