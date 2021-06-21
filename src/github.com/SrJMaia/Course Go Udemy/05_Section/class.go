package main

import (
  "fmt"
  "runtime"
)

var boolVariable bool
var intVariable02 int16
var floatVariable02 float32 // must to be 64 or 32

const constVariable uint16 = 10

const (
  const_one = 1
  const_two = 2
  const_three = 3 // untyped
)

const (
  iotaA = iota
  iotaB
  iotaC
)

const (
  iotaAA = iota
  iotaBB = iota
  iotaCC = iota
)

const (
  //kb = 1024
  //mb = 1024 * kb
  //gb = 1024 * mb
  _ = iota
  kb = 1 << (iota * 10)
  mb = 1 << (iota * 10)
  gb = 1 << (iota * 10)
)

func main()  {

  // Bool Type
  fmt.Println("BOOL TYPES")
  fmt.Println(boolVariable) // false
  boolVariable = true
  fmt.Println(boolVariable)
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
  intVariable01 := 2
  floatVariable01 := .2
  fmt.Println(intVariable01)
  fmt.Printf("%T\n", intVariable01)
  fmt.Println(floatVariable01)
  fmt.Printf("%T\n", floatVariable01)
  intVariable02 = 4
  floatVariable02 = .4
  fmt.Println(intVariable02)
  fmt.Printf("%T\n", intVariable02)
  fmt.Println(floatVariable02)
  fmt.Printf("%T\n", floatVariable02)

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

  stringVariable01 := "Hello World"
  fmt.Printf("%T\n", stringVariable01)
  fmt.Println(stringVariable01)

  slicedOfByte := []byte(stringVariable01)
  fmt.Println(slicedOfByte)
  fmt.Printf("%T\n", slicedOfByte)

  for i := 0; i < len(slicedOfByte); i++ {
    fmt.Printf("%#U ", slicedOfByte[i])
  }

  fmt.Println("")

  for i, v := range slicedOfByte {
    fmt.Println(i, v)
  }

  for i, v := range slicedOfByte {
		fmt.Printf("at index position %d we have hex %#x\n", i, v)
	}

  // Constants
  // const name type = value

  fmt.Println(constVariable)
  fmt.Printf("%T\n", constVariable)

  // Iota

  fmt.Println(iotaA)
  fmt.Println(iotaB)
  fmt.Println(iotaC)
  fmt.Printf("%T\n", iotaA)
  fmt.Printf("%T\n", iotaB)
  fmt.Printf("%T\n", iotaC)

  fmt.Println(iotaAA)
  fmt.Println(iotaBB)
  fmt.Println(iotaCC)
  fmt.Printf("%T\n", iotaAA)
  fmt.Printf("%T\n", iotaBB)
  fmt.Printf("%T\n", iotaCC)

  // Bit Shifting
  bit_shift_01 := 4
  fmt.Printf("%d\t\t%b\n", bit_shift_01, bit_shift_01)

  bit_shift_02 := 4 << 1
  fmt.Printf("%d\t\t%b\n", bit_shift_02, bit_shift_02)

  fmt.Printf("%d\t\t\t%b\n", kb, kb)
  fmt.Printf("%d\t\t\t%b\n", mb, mb)
  fmt.Printf("%d\t\t\t%b\n", gb, gb)



}
