package main

import (
  "fmt"
)

func foo(i... int) int {
  tot := 0
  for _, v := range i {
    tot += v
  }
  return tot
}

func bar(i []int) int {
  tot := 0
  for _, v := range i {
    tot += v
  }
  return tot
}

func main() {
  // Exercise 02
  // • create a func with the identifier foo that
  //     ◦ takes in a variadic parameter of type int
  //     ◦ pass in a value of type []int into your func (unfurl the []int)
  //     ◦ returns the sum of all values of type int passed in
  // • create a func with the identifier bar that
  //     ◦ takes in a parameter of type []int
  //     ◦ returns the sum of all values of type int passed in

  n := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}
  fmt.Println(foo(n...))
  fmt.Println(bar(n))

}
