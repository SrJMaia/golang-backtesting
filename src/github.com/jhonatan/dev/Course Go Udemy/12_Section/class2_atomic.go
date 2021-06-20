package main

import (
  "fmt"
  "runtime"
  "sync"
  "sync/atomic"
)

var wg sync.WaitGroup

func main() {
  // go run -race class2_mutex.go
  fmt.Println("CPUs:\t\t", runtime.NumCPU())
  fmt.Println("Goroutines:\t", runtime.NumGoroutine())

  var counter int64
  const gs = 100
  var wg sync.WaitGroup
  wg.Add(gs)

  for i := 0; i < gs; i++ {
    go func(){
      atomic.AddInt64(&counter, 1) // Add 1 to counter
      fmt.Println("Counter", atomic.LoadInt64(&counter)) // Read counter
      runtime.Gosched()
      wg.Done()
    }()
    fmt.Println("Goroutines:\t", runtime.NumGoroutine())
  }
  wg.Wait()
  fmt.Println("CPUs:\t\t", runtime.NumCPU())
  fmt.Println("Goroutines:\t", runtime.NumGoroutine())
  fmt.Println(counter)
}
