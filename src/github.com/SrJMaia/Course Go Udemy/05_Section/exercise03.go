package main

import (
  "fmt"
)

const (
  typedConst string = "Typed Hello World"
  untypedConst = "Untyped Hello World"
)

func main() {
  // Exercise 03
  // Create TYPED and UNTYPED constants. Print the values of the constants.

  fmt.Println(typedConst)
  fmt.Println(untypedConst)
}
