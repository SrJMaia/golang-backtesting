package main

import (
  "fmt"
)

func main() {
  // Exercise 08
  // Create a program that uses a switch statement with no switch expression specified.

  switch {
	case false:
		fmt.Println("this should not print")
	case true:
		fmt.Println("this should print")
  }

}
