package main

import (
  "fmt"
)

func main() {
  // Exercise 04
  // Create and use an anonymous struct

  candle := struct {
    open float32
    high float32
    low float32
    close float32
  }{
    open: 1.12230,
    high: 1.12401,
    low: 1.12200,
    close: 1.12210,
  }

  fmt.Println(candle)
}
