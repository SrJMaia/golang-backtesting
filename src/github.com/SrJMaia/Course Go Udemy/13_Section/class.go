package main

import (
  "fmt"
)

func main() {
  // chanels
  var c = make(chan int)

  go func(){
      c <- 42
    }()

  fmt.Println(<- c)

  var cc = make(chan int, 1) // Buffer channel

  cc <- 42

  fmt.Println(<- cc)

  var ccc = make(chan int, 2) // Buffer channel

  ccc <- 42
  ccc <- 43

  fmt.Println(<- ccc)
  fmt.Println(<- ccc)
  fmt.Println("-------")
  fmt.Printf("%T\n", c)

  var ci = make(chan int)
  var cii = make(<-chan int) // receive
  var ciii = make(chan<- int) // send

  fmt.Println("-------")
  fmt.Printf("%T\n", ci)
  fmt.Printf("%T\n", cii)
  fmt.Printf("%T\n", ciii)
  fmt.Println("-------")
  //fmt.Printf("%T\n", (chan int)(cii))
  //fmt.Printf("%T\n", (chan int)(ciii))

  // odes not run
  // var c = make(<-chan int)
  // go func(){
  //   c <- 42
  //   }()
  // fmt.Println(<-c)
  // fmt.Println("-------")
  // fmt.Printf("%T\n", c)



}
