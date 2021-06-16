package main

import (
  "fmt"
)

func main() {
  // Exercise 01
  // • Using a COMPOSITE LITERAL:
  // • create an ARRAY which holds 5 VALUES of TYPE int
  // • assign VALUES to each index position.
  // • Range over the array and print the values out.
  // • Using format printing
  //     ◦ print out the TYPE of the array

  // Mine
  // var myArray [5]int
  // for j := 0; j < 5; j++ {
  //   myArray[j] = j*10
  // }

  // Teacher
  var myArray [5]int{0, 1, 2, 3, 4}
  for _, v := range myArray {
    fmt.Println(v)
  }

  fmt.Printf("%T", myArray)

}
