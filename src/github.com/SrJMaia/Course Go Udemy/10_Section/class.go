package main

import (
  "fmt"
)

func main() {
  // Pointers
  var a int = 42
  fmt.Println(a)
  fmt.Println(&a) // Memory Address
  fmt.Printf("%T\n", a)
  fmt.Printf("%T\n", &a) // Its a pointer to an int, its a different type

  //var b *int = &a
  b := &a // Here i have the address
  fmt.Println(b)
  fmt.Println(*b) // Here i print the value in that address
  fmt.Println(*&b)

  *b = 43
  fmt.Println(*b)
  fmt.Println(a)

  // Its used to pass data, instead to dup the data, you point the location
  var x int = 0
  fmt.Println("x befor", &x)
  fmt.Println("x befor", x)
  foo(&x)
  fmt.Println("x after", &x)
  fmt.Println("x after", x)

  // Methods sets
  
}

func foo(y *int) {
  fmt.Println("y befor", &y)
  fmt.Println("y befor", y)
  *y = 43
  fmt.Println("y after", &y)
  fmt.Println("y after", y)
}
