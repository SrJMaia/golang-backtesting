package main

import (
  "fmt"
)

func main() {
  // Exercise 07
  // • write a program that
  //     ◦ launches 10 goroutines
  //         ▪ each goroutine adds 10 numbers to a channel
  //     ◦ pull the numbers off the channel and print them
  //solutions:
  //     • https://play.golang.org/p/R-zqsKS03P
  //     • https://play.golang.org/p/quWnlwzs2z
  // student solutions:
  //     • https://twitter.com/mannion
  //         ◦ https://gist.github.com/mannion007/3c8899913974c1027ef6f13ec37b2b3f
  c := make(chan int)

	for j := 0; j < 10; j++ {
		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
		}()
	}

	for k := 0; k < 100; k++ {
		fmt.Println(k, <-c)
	}

	fmt.Println("about to exit")
}
