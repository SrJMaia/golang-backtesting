package main

import (
  "fmt"
  "time"
)

func main() {
  // Exercise 03
  // Create a for loop using this syntax
  //     • for condition { }
  // Have it print out the years you have been alive.

  i := 1998
  for i <= time.Now().Year() {
    fmt.Println("Year: ", i)
    i++
  }
}
