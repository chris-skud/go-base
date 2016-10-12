package main

import (
	"fmt"
	"time"
)

// demonstrate the use of a semaphore to throttle concurrent outbound requests
func main() {

	// semaphore to limit concurent requests
	const maxConcurrent = 10
	sem := make(chan int, maxConcurrent)

	boring(sem)

	// run until interrupt
	for {
	}
}

func boring(sem chan int) {
	go func() {
		for i := 0; ; i++ {
			sem <- 1
			do(i)
			<-sem
		}
	}()
}

func do(i int) {
	time.Sleep(1 * time.Second)
	fmt.Println(i)
}
