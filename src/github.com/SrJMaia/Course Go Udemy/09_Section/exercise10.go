package main

import (
  "fmt"
)

func main() {
  // Exercise 10
  // Closure is when we have “enclosed” the scope of a variable in some code block.
  // For this hands-on exercise, create a func which “encloses” the scope of a variable:

  any()
}

func any () {
  a := "Far"
  fmt.Println(a)
  {
    b := "Close"
    fmt.Println(b)
  }
}
