package main

import (
  "fmt"
  "runtime"
)

func main() {
  // Exercise 06
  // Create a program that prints out your OS and ARCH. Use the following commands to run it
  //     • go run
  //     • go build
  //     • go install
  fmt.Println(runtime.GOOS, runtime.GOARCH)

}
