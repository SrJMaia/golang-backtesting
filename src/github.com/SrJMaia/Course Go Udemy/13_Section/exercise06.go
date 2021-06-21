package main

import (
  "fmt"
)

func main() {
  // Exercise 06
  // • write a program that
  //     ◦ puts 100 numbers to a channel
  //     ◦ pull the numbers off the channel and print them
  c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}
