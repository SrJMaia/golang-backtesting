package main

import (
  "fmt"
)

func main() {
  // Exercise 09
  // Create a program that uses a switch statement with the switch
  // expression specified as a variable of TYPE string with the IDENTIFIER “favSport”.
  favSport := "Something"
  switch favSport {
  case "A":
    fmt.Println("this should not print")
  case "B":
    fmt.Println("this should not print")
  default:
    fmt.Println("this should print")
  }


}
