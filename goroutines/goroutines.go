package main

import (
	"fmt"
	"time"
)

func main() {
	// count("foo")
	// count("bar")

	go count("foo")
	go count("bar")
	time.Sleep(6 * time.Second)
}

func count(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%s: %d\n", name, i)
		time.Sleep(500 * time.Millisecond)
	}
}
