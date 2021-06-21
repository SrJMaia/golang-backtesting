package main

import (
  "fmt"
)

func foo() int {
  return 10
}

func bar() (int, string) {
  return 10, "ten"
}

func main() {
  // Exercise 01
  // create a func with the identifier foo that returns an int
  // create a func with the identifier bar that returns an int and a string
  // call both funcs
  // print out their results
  fmt.Println(foo())
  fmt.Println(bar())

}
