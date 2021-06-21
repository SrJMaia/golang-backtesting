package main

import (
  "fmt"
  "time"
)

func main() {
  // Exercise 04
  // Create a for loop using this syntax
  //     • for { }
  // Have it print out the years you have been alive.

  i := 1998
  for {
    if i > time.Now().Year() {
      break
    }
    fmt.Println("Year: ", i)
    i++
  }
}
