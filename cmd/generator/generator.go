package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Generator pattern (Rob Pike: https://www.youtube.com/watch?v=f6kdp27TYZs)
// is a function that returns a channel
func main() {
	c := boring("hey")
	c2 := boring("you")
	for i := 0; i < 5; i++ {
		fmt.Println(<-c)
		fmt.Println(<-c2)
	}

	fmt.Println("exit")
}

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			fmt.Println("doing", i)
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	return c
}
