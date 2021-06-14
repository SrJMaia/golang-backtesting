package main

import (
  "fmt"
)

func main() {
  // Exercise 04
  // Write a program that
  //     • assigns an int to a variable
  //     • prints that int in decimal, binary, and hex
  //     • shifts the bits of that int over 1 position to the left, and assigns that to a variable
  //     • prints that variable in decimal, binary, and hex

  intValue := 100
  fmt.Printf("%d\t%b\t%#x\n", intValue, intValue, intValue)

  shiftedIntValue := intValue << 1
  fmt.Printf("%d\t%b\t%#x", shiftedIntValue, shiftedIntValue, shiftedIntValue)

}
