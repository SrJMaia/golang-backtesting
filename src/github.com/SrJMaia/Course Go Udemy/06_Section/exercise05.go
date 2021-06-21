package main

import (
  "fmt"
)

func main() {
  // Exercise 05
  // Print out the remainder (modulus) which is found for
  // each number between 10 and 100 when it is divided by 4
  for i := 10; i <= 100; i++ {
    fmt.Println("Remainder: ", i%4)
  }

}
