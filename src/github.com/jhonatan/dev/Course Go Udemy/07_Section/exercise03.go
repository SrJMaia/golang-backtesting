package main

import (
  "fmt"
)

func main() {
  // Exercise 03
  // Using the code from the previous example, use SLICING to create the
  // following new slices which are then printed:
  //     • [42 43 44 45 46]
  //     • [47 48 49 50 51]
  //     • [44 45 46 47 48]
  //     • [43 44 45 46 47]

  slice01 := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

  fmt.Println(slice01[:5])
  fmt.Println(slice01[5:])
  fmt.Println(slice01[2:7])
  fmt.Println(slice01[1:6])
}
