package main

import "fmt"

func main()  {
  number_bytes, erros := fmt.Println("Teste", 42, true)
  fmt.Println(number_bytes)
  fmt.Println(erros)
  teste, _ := fmt.Println("Teste", 42, true)
  fmt.Println(teste)
  // You can't ignore a return, use _ to ignore or nothing at all
  // interface{} paramater in the doc is the same as args** and kwargs**


  // ---------------- Variables
  // Identifiers are the variables names
  // identifier = letter {letter or | unicode_digit}
  x := 42 // Declare and assigning
  fmt.Println(x)
  x = 10 // Just assig
  fmt.Println(x)
  name := "Something"
  fmt.Println(name)
}
