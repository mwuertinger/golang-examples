package main

import (
	"fmt"
	"time"
)

func main() {
	foo := make(chan int, 0)
	bar := make(chan int, 0)

	go work(50, foo)
	go work(150, bar)

	tick := time.Tick(1000 * time.Millisecond)

	totalProgress := 0
	p1 := 0
	p2 := 0
	for totalProgress < 200 {
		select {
		case s := <-foo:
			totalProgress += s
			p1 += s
		case s := <-bar:
			totalProgress += s
			p2 += s
		case <-tick:
			fmt.Print("\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r")
			fmt.Printf("Progress: %v (%v, %v)", totalProgress, p1, p2)
		}
	}
}

func work(sleepDuration int, progress chan int) {
	for i := 0; i < 100; i++ {
		progress <- 1
		time.Sleep(time.Duration(sleepDuration) * time.Millisecond)
	}
}
