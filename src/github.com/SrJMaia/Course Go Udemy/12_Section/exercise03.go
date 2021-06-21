package main

import (
  "fmt"
  "sync"
  "runtime"
)

func main() {
  // Exercise 03
  // • Using goroutines, create an incrementer program
  //     ◦ have a variable to hold the incrementer value
  //     ◦ launch a bunch of goroutines
  //         ▪ each goroutine should
  //             • read the incrementer value
  //                 ◦ store it in a new variable
  //             • yield the processor with runtime.Gosched()
  //             • increment the new variable
  //             • write the value in the new variable back to the incrementer variable
  // • use waitgroups to wait for all of your goroutines to finish
  // • the above will create a race condition.
  // • Prove that it is a race condition by using the -race flag
  // • if you need help, here is a hint: https://play.golang.org/p/FYGoflKQej
  var wg sync.WaitGroup

  var incrementer int = 0
  var gs int = 100
  wg.Add(gs)

  for i := 0; i < gs; i++ {
    go func(){
      v := incrementer
      runtime.Gosched()
      v++
      incrementer = v
      fmt.Println(incrementer)
      wg.Done()
    }()
   }
   wg.Wait()
   fmt.Println(incrementer)
}
