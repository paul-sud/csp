/*
Problem: Construct a buffering process X to smooth
variations in the speed of output of portions by a
producer process and input by a consumer process. The
consumer contains pairs of commands X!more( );
X?p, and the producer contains commands of the form
X!p. The buffer should contain up to ten portions.

Note: In the classic solution we would need a separate goroutine sitting in between the
producer and consumer, but in Go we can use buffered channels for this purpose. Here we
actually use a waitgroup instead of sleeping since it could take a while for all of the
values to get sent.
*/
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	buf := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go producer(buf)
	}
	go consumer(buf, &wg)
	wg.Wait()
}

func producer(buf chan<- bool) {
	buf <- true
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
}

func consumer(buf <-chan bool, wg *sync.WaitGroup) {
	for val := range buf {
		wg.Done()
		fmt.Println("Received val", val)
		time.Sleep(50 * time.Millisecond)
	}
}
