package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  // Exercise 07
  // Building on the previous hands-on exercise, create a program that uses “else if” and “else”.

  rand.Seed(time.Now().UnixNano())
  x := rand.Intn(3)

  if x == 0 {
    fmt.Println("Scissors")
  } else if x == 1 {
    fmt.Println("Rock")
  } else {
    fmt.Println("Paper")
  }

}
