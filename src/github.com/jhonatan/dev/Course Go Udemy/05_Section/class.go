package main

import (
  "fmt"
  "runtime"
)

var bool_variable bool
var int_variable02 int16
var float_variable02 float32 // must to be 64 or 32

func main()  {

  // Bool Type
  fmt.Println("BOOL TYPES")
  fmt.Println(bool_variable) // false
  bool_variable = true
  fmt.Println(bool_variable)
  comparassion01 := 10
  comparassion02 := 20
  fmt.Println(comparassion01, comparassion02 )
  fmt.Println("Equal: ", comparassion01 == comparassion02)
  fmt.Println("Greater: ", comparassion01 > comparassion02)
  fmt.Println("Greater or Equal: ", comparassion01 >= comparassion02)
  fmt.Println("Less or Equal: ", comparassion01 <= comparassion02)
  fmt.Println("Less: ", comparassion01 < comparassion02)
  fmt.Println("Different: ", comparassion01 != comparassion02)

  // Numerical Types
  fmt.Println("NUMERICAL TYPES")
  int_variable01 := 2
  float_variable01 := .2
  fmt.Println(int_variable01)
  fmt.Printf("%T\n", int_variable01)
  fmt.Println(float_variable01)
  fmt.Printf("%T\n", float_variable01)
  int_variable02 = 4
  float_variable02 = .4
  fmt.Println(int_variable02)
  fmt.Printf("%T\n", int_variable02)
  fmt.Println(float_variable02)
  fmt.Printf("%T\n", float_variable02)

  // uint8       the set of all unsigned  8-bit integers (0 to 255)
  // uint16      the set of all unsigned 16-bit integers (0 to 65535)
  // uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
  // uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)
  //
  // int8        the set of all signed  8-bit integers (-128 to 127)
  // int16       the set of all signed 16-bit integers (-32768 to 32767)
  // int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
  // int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)
  //
  // float32     the set of all IEEE-754 32-bit floating-point numbers
  // float64     the set of all IEEE-754 64-bit floating-point numbers
  //
  // complex64   the set of all complex numbers with float32 real and imaginary parts
  // complex128  the set of all complex numbers with float64 real and imaginary parts
  //
  // byte        alias for uint8
  // rune        alias for int32

  // FUN FACT
  fmt.Println(runtime.GOOS)
  fmt.Println(runtime.GOARCH)

  // String Type

  string_variable01 := "Hello World"
  fmt.Printf("%T\n", string_variable01)
  fmt.Println(string_variable01)

  sliced_of_byte := []byte(string_variable01)
  fmt.Println(sliced_of_byte)
  fmt.Printf("%T\n", sliced_of_byte)

  for i := 0; i < len(sliced_of_byte); i++ {
    fmt.Printf("%#U ", sliced_of_byte[i])
  }

  fmt.Println("")

  for i, v := range sliced_of_byte {
    fmt.Println(i, v)
  }

  for i, v := range sliced_of_byte {
		fmt.Printf("at index position %d we have hex %#x\n", i, v)
	}



}
