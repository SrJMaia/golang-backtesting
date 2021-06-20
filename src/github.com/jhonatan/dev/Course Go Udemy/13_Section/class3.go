package main

import (
	"fmt"
)

func main() {
	// chanels
	var c = make(chan int)

	// send
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()
	// receive
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("about to exit")

}
