package main

import (
  "fmt"
  "runtime"
  "sync"
)

var wg sync.WaitGroup

func main() {
  // go run -race class2_mutex.go
  fmt.Println("CPUs:\t\t", runtime.NumCPU())
  fmt.Println("Goroutines:\t", runtime.NumGoroutine())

  var counter int = 10
  const gs = 100
  var wg sync.WaitGroup
  wg.Add(gs)

  var mu sync.Mutex
  for i := 0; i < gs; i++ {
    go func(){
      mu.Lock()
      v := counter
      runtime.Gosched()
      v ++
      counter = v
      mu.Unlock()
      wg.Done()
    }()
    fmt.Println("Goroutines:\t", runtime.NumGoroutine())
  }
  wg.Wait()
  fmt.Println("CPUs:\t\t", runtime.NumCPU())
  fmt.Println("Goroutines:\t", runtime.NumGoroutine())
  fmt.Println(counter)
}
