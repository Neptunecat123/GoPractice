package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 5; i++ {
		fmt.Println(from, ":", i)
	}
}

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	f("main")

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")

	// default channel
	message := make(chan string)
	go func() {
		message <- "ping"
	}()

	m := <-message
	fmt.Println(m)

	// buffer channel
	message = make(chan string, 2)
	message <- "buffer"
	message <- "channel"

	fmt.Println(<-message)
	fmt.Println(<-message)

	// channel sync
	done := make(chan bool, 1)
	go worker(done)

	<-done

	// direction
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed massage")
	pong(pings, pongs)
	fmt.Println(<-pongs)

	// select
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		// time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		// time.Sleep(1 * time.Second)
		c1 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		default:
			fmt.Println("no activity")
		}
	}
	close(c1)
	fmt.Println("close c1")
}
