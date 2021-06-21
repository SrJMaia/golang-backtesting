package main

import (
  "fmt"
  "time"
)

func main() {
  // Exercise 06
  // Build and use an anonymous func
  start := time.Now()
  func (x int) {
    fmt.Println("Fibonacci Sequence")
    prev, pos := 0, 1
    for ; x != 0; x-- {
      pos += prev
      prev = pos - prev
      }
    fmt.Println(prev)
  }(100)

  fmt.Print(time.Since(start))


}
