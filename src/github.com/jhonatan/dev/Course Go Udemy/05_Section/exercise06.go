package main

import (
  "fmt"
)

const (
  iotaA = iota + 2021
  iotaB
  iotaC
)

func main() {
  // Exercise 06
  // Using iota, create 4 constants for the NEXT 4 years. Print the constant values.

  fmt.Println(iotaA)
  fmt.Println(iotaB)
  fmt.Println(iotaC)
}
