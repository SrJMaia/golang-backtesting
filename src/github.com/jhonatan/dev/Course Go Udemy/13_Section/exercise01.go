package main

import (
	"fmt"
)

func main() {
  // Exercise  01
  // get this code working:
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)

  ci := make(chan int, 1)

	ci <- 42

	fmt.Println(<-ci)
}
