package main

import "fmt"

// if you declare a variable outside a function uses var
// when use var, you dont use :=
var outsideFunc = 10
var iOutside int // is the same as var iOutside int = 0
// The value of 0 will be definied to int, for float will be 0.0
// bool == false, string == ""
var text string = "Someone said: 'teste12'"

var a int
type hotdog int
var b hotdog

func main()  {
  numberBytes, erros := fmt.Println("Teste", 42, true)
  fmt.Println(numberBytes)
  fmt.Println(erros)
  teste, _ := fmt.Println("Teste", 42, true)
  fmt.Println(teste)
  // You can't ignore a return, use _ to ignore or nothing at all
  // interface{} paramater in the doc is the same as args** and kwargs**

  // ---------------- Variables Short Declaration
  // Identifiers are the variables names
  // identifier = letter {letter or | unicode_digit}
  x := 42 // Declare and assigning
  fmt.Println(x)
  x = 10 // Just assig
  fmt.Println(x)
  name := "Something"
  fmt.Println(name)

  // ---------------- var keyword for variable declaration
  // Use when its outside
  var x1 = 10
  fmt.Println(x1)

  fmt.Println(text)
  fmt.Printf("%T\n", text) // To see the type


  // ---------------- fmt package
  fmt.Println(outsideFunc)
  fmt.Printf("%T\n", outsideFunc) // type
  fmt.Printf("%b\n", outsideFunc) // binary
  fmt.Printf("%x\n", outsideFunc) // hexadecimal
  fmt.Printf("%#x\n", outsideFunc) // hexadecimal with 0x
  fmt.Printf("%#x\t%b\t%x", outsideFunc, outsideFunc, outsideFunc) // hexadecimal with 0x, binary, hexadecimal


  // ---------------- creating your own type
  // the code below should be outside func main
  // var a int
  // type hotdog int
  // var b hotdog
  // We create a type hotdog that is an int underline, or have int underline
  // a = b but we cant do that below
  // to use as an int we do this
  a = int(b)
  // Thats not casting, but its conversion

}
