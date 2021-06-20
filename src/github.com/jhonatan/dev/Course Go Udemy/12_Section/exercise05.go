package main

import (
  "fmt"
  "sync"
  "sync/atomic"
)

func main() {
  // Exercise 05
  // Fix the race condition you created in exercise #4 by using package atomic
  var wg sync.WaitGroup

  var incrementer int64
  var gs int = 100
  wg.Add(gs)

  for i := 0; i < gs; i++ {
    go func(){
      atomic.AddInt64(&incrementer, 1)
      fmt.Println(atomic.LoadInt64(&incrementer))
      wg.Done()
    }()
   }
   wg.Wait()
   fmt.Println(incrementer)

}
