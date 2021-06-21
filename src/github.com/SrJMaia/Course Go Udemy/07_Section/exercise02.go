package main

import (
  "fmt"
)

func main() {
  // Exercise 02
  // • Using a COMPOSITE LITERAL:
  // • create a SLICE of TYPE int
  // • assign 10 VALUES
  // • Range over the slice and print the values out.
  // • Using format printing
  //     ◦ print out the TYPE of the slice

  mySlice := []int{0, 10, 20, 30, 40, 50, 60, 70, 80, 90}

  for _, v := range mySlice {
    fmt.Println(v)
  }

  fmt.Printf("%T", mySlice)

}
