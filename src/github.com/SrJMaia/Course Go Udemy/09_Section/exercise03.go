package main

import (
  "fmt"
)

func first() {
  fmt.Println("First")
}

func second() {
  fmt.Println("Second")
}

func main() {
  // Exercise 03
  // Use the “defer” keyword to show that a deferred func runs
  // after the func containing it exits.
  defer first()
  second()
}
