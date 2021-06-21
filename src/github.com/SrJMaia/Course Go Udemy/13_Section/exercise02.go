package main

import (
	"fmt"
)

func main() {
  // Exercise 02
  // • Get this code running:
  //     ◦ https://play.golang.org/p/oB-p3KMiH6
  //         ▪ solution: https://play.golang.org/p/isnJ8hMMKg
  //     ◦ https://play.golang.org/p/_DBRueImEq
  //         ▪ solution: https://play.golang.org/p/mgw750EPp4
	cs := make(chan int)

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)
.
	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)

  cr := make(chan int)

	go func() {
		cr <- 42
	}()
	fmt.Println(<-cr)

	fmt.Printf("------\n")
	fmt.Printf("cr\t%T\n", cr)
}
