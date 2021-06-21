package main

import (
  "fmt"
)

func main() {
  // Exercise 09
  // A “callback” is when we pass a func into a func as an argument.
  // For this exercise,
  //     • pass a func into a func as an argument
  a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}

  fmt.Println(sum(a...))
  fmt.Print(even(sum, a...))



}

func sum(x1 ...int) int {
  tot := 0
  for _, v := range x1 {
    tot += v
  }
  return tot
}

func even(f func(xi ...int) int, vi ...int) int {
  var yi []int
  for _, v := range vi {
    if v % 2 == 0 {
      yi = append(yi, v)
    }
  }
  return f(yi...)
}
