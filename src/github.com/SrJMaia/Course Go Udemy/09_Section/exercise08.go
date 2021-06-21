package main

import (
  "fmt"
)

func main() {
  // Exercise 08
  // • Create a func which returns a func
  // • assign the returned func to a variable
  // • call the returned func

  teste := first()
  fmt.Println(teste())


}

func first() func() int {
  return func() int {
    return 10
  }
}
