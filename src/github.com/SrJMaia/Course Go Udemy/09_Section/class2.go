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

func (s secretAgent) speak() {
  fmt.Println("I am ", s.first, s.last, " - the secretAgent speak")
}

func (p person) speak() {
  fmt.Println("I am ", p.first, p.last, " - the person speak")
}

type human interface {
  speak()
}

func bar(h human) {
  switch h.(type) {
  case person:
    fmt.Println("I'm a person", h.(person).first)
  case secretAgent:
    fmt.Println("I'm a secretAgent", h.(secretAgent).first)
  }
  fmt.Println("I was passed into bar", h)
}

// ***
// func SwitchOnType(x interface{}) {
//   switch x.(type) {
//   case int:
//     fmt.Println("int")
//   case string:
//     fmt.Println("string")
//   case concat:
//     fmt.Println("concat")
//   default:
//     fmt.Println("unknown")
//   }
// }

func foo() string {
  s := "Hello World"
  return s
}

func bar1() func() int {
  return func() int  {
    return 451
  }
}

func sum(x1 ...int) int {
  tot := 0
  for _, v := range x1 {
    tot += v
  }
  return tot
}

// f will be a func(receive variadic paramaters) return an int
func even(f func(xi ...int) int, vi ...int) int {
  var yi []int
  for _, v := range vi {
    if v % 2 == 0 {
      yi = append(yi, v)
    }
  }
  return f(yi...)
}

func odd(f func(xi ...int) int, vi ...int) int {
  var yi []int
  for _, v := range vi {
    if v % 2 != 0 {
      yi = append(yi, v)
    }
  }
  return f(yi...)
}

func factorial(n int) int {
  if n == 0 {
    return 1
  }
  return n * factorial(n-1)
}

func BetterFactorial(n int) int {
  result := 1
  for ; n != 0; n-- {
    result *= n
  }
  return result
}

func main() {
  // Interfaces & Polymorphism
  // p1 is a type secretAgent, because we have type human intercae, he is also
  // a type human. A value can have different types
  sa1 := secretAgent {
    person: person {
      first: "James",
      last: "Bond",
    },
    ltk:true,
  }

  fmt.Println(sa1)
  sa1.speak()

  p1 := person {
    first: "Dr.",
    last: "Yes",
  }

  fmt.Println(p1)
  bar(sa1)
  bar(p1)

  // Conversion ***

  // Anonymous Functions

  func() {
    fmt.Println("Anonymous func ran")
  }() // The final paratensis execute the func

  func(x int) {
    fmt.Println("Anonymous func ran: ",x)
  }(42)

  // Func Expression | Functions is a type like int or float
  f := func() {
    fmt.Println("my first func expression")
  }
  f()
  fmt.Printf("%T\n", f)

  // Returning a func

  s1 := foo()
  fmt.Println(s1)
  b1 := bar1()
  fmt.Printf("%T\n", b1)
  rb1 := b1()
  fmt.Println(rb1)
  fmt.Println(b1())
  fmt.Println(bar1()())

  // Callback | Using fucntions as paramater and argument

  ii := []int{1,2,3,4,5,6,7,7,8,9,}
  sum1 := sum(ii...)
  fmt.Println(sum1)

  sum2 := even(sum, ii...)
  fmt.Println(sum2)
  sum3 := odd(sum, ii...)
  fmt.Println(sum3)

  // Closure
  {
    // y will only work inside this block
    y := 42
    fmt.Println(y)
  }

  // Recursion
  // When a function call itself
  fac := factorial(4)
  fmt.Println(fac)
  fmt.Println(BetterFactorial(4))

}
