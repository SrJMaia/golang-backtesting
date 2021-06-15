package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  // Exercise 06
  // Create a program that shows the “if statement” in action.
  rand.Seed(time.Now().UnixNano())
  x := rand.Intn(3)

  if x == 0 {
    fmt.Println("Scissors")
  }


}
