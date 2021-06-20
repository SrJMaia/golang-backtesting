package main

import (
  "fmt"
  "runtime"
  "sync"
)

var wg sync.WaitGroup

func main() {
  fmt.Println("CPUs:\t\t", runtime.NumCPU())
  fmt.Println("Goroutines:\t", runtime.NumGoroutine())

  var counter int = 10
  const gs = 100
  var wg sync.WaitGroup
  wg.Add(gs)
  for i := 0; i < gs; i++ {
    go func(){
      v := counter
      runtime.Gosched()
      v ++
      counter = v
      wg.Done()
    }()
    fmt.Println("Goroutines:\t", runtime.NumGoroutine())
  }
  wg.Wait()
  fmt.Println("CPUs:\t\t", runtime.NumCPU())
  fmt.Println("Goroutines:\t", runtime.NumGoroutine())
  fmt.Println(counter)
}
