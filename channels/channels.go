package main

import (
	"fmt"
	"time"
)

func main() {
	doneFoo := make(chan bool)
	doneBar := make(chan bool)

	go count("foo", doneFoo)
	go count("bar", doneBar)

	<-doneFoo
	<-doneBar
}

func count(name string, done chan bool) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%s: %d\n", name, i)
		time.Sleep(500 * time.Millisecond)
	}
	done <- true
}
