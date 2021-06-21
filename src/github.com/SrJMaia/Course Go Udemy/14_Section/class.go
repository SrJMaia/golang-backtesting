package main

import (
  "fmt"
  "os"
  "log"
)

func main() {
  // errors
	var _, err = fmt.Println("Teste")
  if err != nil {
    panic(err)
  }
  if err != nil {
    fmt.Println(err)
    return
  }

  // _, err = os.Open("no-file.txt")
  // if err != nil {
  //   fmt.Println("err happened", err)
  // }

  // Maybe the "best" for me
  // _, err = os.Open("no-file.txt")
  // if err != nil {
  //   log.Println("err happened", err) // with log prints the date and hour
  // }

  // _, err = os.Open("no-file.txt")
  // if err != nil {
  //   log.Fatalln("err happened", err) // Return status != 0 if its an error
  // }

  _, err = os.Open("no-file.txt")
  if err != nil {
    log.Panicln("err happened", err) // Its the same as panic
  }

}
