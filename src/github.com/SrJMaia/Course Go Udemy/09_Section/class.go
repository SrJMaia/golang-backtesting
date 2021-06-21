package main

import (
  "fmt"
)

type person struct {
  first string
  last string
}

type secretAgent struct {
  person
  ltk bool
}

func main() {
  // Functions
  foo()
  bar("This is an argument")
  result := mySum(10, 20)
  fmt.Println(result)
  fmt.Println(mySumTwo(20, 30))
  fmt.Println(multipleReturns(20, 30))

  // Variadic Paramaters 0 or more
  // Must be the final paramater
  multiplePara(0, 1, 2, 3, 4, 5)

  // Unfurling a slice
  x := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
  multiplePara(x...)
  multiplePara()

  // Defer - Basically forces the function to run the last
  // PS: It will run, but everything will be stored, and than when the surrounding
  // functions finish go will output the return
  defer foo()
  bar("Teste")

  // Methods

  p1 := secretAgent {
    person: person {
      first: "James",
      last: "Bond",
    },
    ltk:true,
  }

  fmt.Println(p1)
  p1.speak()
}

func foo() {
  fmt.Println("Hello World")
}

// s here is a paramater
func bar(s string) {
  fmt.Println(s)
}

func mySum(x int, y int) int {
  return x+y
}

func mySumTwo(x, y int) int { // You can do this way
  return x+y
}

func multipleReturns(x, y int) (int, bool) { // You can do this way
  return x+y, true
}

func multiplePara(x ...int) {
  fmt.Println(x)
  fmt.Printf("%T\n", x)

  sum := 0
  for k, v := range x {
    fmt.Println("Key: ", k, " Sum: ", sum)
    sum += v
  }
}

// Any value of secretAgent type will have this method
func (s secretAgent) speak() {
  fmt.Println("I am ", s.first, s.last)
}
