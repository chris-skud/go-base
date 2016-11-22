package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(<-chan string)
	c = fanIn(boring("Bill"), boring("Jane"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}

	fmt.Println("boring, i'm leaving")
}

func boring(msg string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(2e3)) * time.Millisecond)
		}
	}()

	return c
}

func fanIn(ch1, ch2 <-chan string) <-chan string {
	ch := make(chan string)
	go func() {
		for {
			// reads as: write string from channel ch1 to channel ch
			ch <- <-ch1
		}
	}()
	go func() {
		for {
			ch <- <-ch2
		}
	}()
	return ch
}
