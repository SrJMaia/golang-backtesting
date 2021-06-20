package main

import (
  "fmt"
  "sync"
)

var wg sync.WaitGroup

func main() {
  // Exercise 01
  // • in addition to the main goroutine, launch two additional goroutines
  //     ◦ each additional goroutine should print something out
  // • use waitgroups to make sure each goroutine finishes before your program exists
  go p1()
  go p2()
  p3()

  wg.Wait()


}

func p1() {
  fmt.Println("Go Routine 1")
}

func p2() {
  fmt.Println("Go Routine 2")
}

func p3() {
  fmt.Println("Go Routine 3")
}
