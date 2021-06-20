package main

import (
  "fmt"
)

type user struct {
  first string
  last string
  age uint8
}

func (s user) speak() {
  fmt.Println("I am ", s.first, "and I am ", s.age, "years old.")
}

func main() {
  // Exercise 04
  // • Create a user defined struct with
  //     ◦ the identifier “person”
  //     ◦ the fields:
  //         ▪ first
  //         ▪ last
  //         ▪ age
  // • attach a method to type person with
  //     ◦ the identifier “speak”
  //     ◦ the method should have the person say their name and age
  // • create a value of type person
  // • call the method from the value of type person

  person := user{
    first:"James",
    last:"Bond",
    age:32,
  }

  person.speak()





}
