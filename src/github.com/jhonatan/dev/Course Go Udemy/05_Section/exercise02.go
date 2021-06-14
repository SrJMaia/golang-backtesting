package main

import (
  "fmt"
)

func main() {
  // Exercise 02
  // Using the following operators, write expressions and assign their values to variables:
  //   g. ==
  //   h. <=
  //   i. >=
  //   j. !=
  //   k. <
  //   l. >
  // Now print each of the variables.

  a := 42
  equal := (a == a)
  equalLess := (a <= a)
  equalGreater := (a >= a)
  different := (a != a)
  less := (a < a)
  greater := (a > a)

  fmt.Println(equal)
  fmt.Println(equalLess)
  fmt.Println(equalGreater)
  fmt.Println(different)
  fmt.Println(less)
  fmt.Println(greater)
}
