package main

import (
  "fmt"
)


type customErr struct {
	info string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("here is the error: %v", ce.info)
}

func main() {
  // Exercise 02
  // Create a struct “customErr” which implements the builtin.error interface.
  // Create a func “foo” that has a value of type error as a parameter.
  // Create a value of type “customErr” and pass it into “foo”. If you need a hint, here is one.
  c1 := customErr{
		info: "need more coffee",
	}

	foo(c1)
}

func foo(e error) {
	fmt.Println("foo ran -", e, "\n")
}
