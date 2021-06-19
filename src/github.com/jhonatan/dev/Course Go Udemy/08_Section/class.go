package main

import (
  "fmt"
)

type person struct {
  first string
  last string
  age int
}

type secretAgent struct {
  person
  ltk bool
}

func main() {
  // Struct similar to a class
  // Always go to large to least. First int64 float32....
  p1 := person {
    first: "James",
    last: "Bond",
    age: 30,
  }

  p2 := person {
    first: "Miss",
    last: "Moneypenny",
    age: 32,
  }

  fmt.Println(p1)
  fmt.Println(p2)
  fmt.Println(p1.first)
  fmt.Println(p2.first)
  fmt.Println(p1.last)
  fmt.Println(p2.last)
  fmt.Println(p1.age)
  fmt.Println(p2.age)

  sap1 := secretAgent {
    person: person {
      first: "James",
      last: "Bond",
      age: 30,
    },
    ltk: true,
  }

  sap2 := secretAgent {
    person: person {
      first: "Miss",
      last: "Moneypenny",
      age: 32,
    },
    ltk: true,
  }

  fmt.Println(sap1.first)
  fmt.Println(sap1.last)
  fmt.Println(sap1.person.last) // Avoid colision
  fmt.Println(sap1.age)
  fmt.Println(sap1.ltk)
  fmt.Println(sap2.first)
  fmt.Println(sap2.last)
  fmt.Println(sap2.age)
  fmt.Println(sap2.ltk)

  // Anonymous Struct
  p3 := struct {
    first string
    last string
    age int
  }{
    first: "James",
    last: "Bond",
    age: 30,
  }

  fmt.Println(p3)

}
