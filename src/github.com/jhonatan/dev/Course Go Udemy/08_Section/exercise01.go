package main

import (
  "fmt"
)

type person struct {
  firstName string
  lastName string
  favoriteIceCream []string
}

func main() {
  // Exercise 01
  // Create your own type “person” which will have an underlying
  // type of “struct” so that it can store the following data:
  //     • first name
  //     • last name
  //     • favorite ice cream flavors
  // Create two VALUES of TYPE person. Print out the values,
  // ranging over the elements in the slice which stores the favorite flavors.

  p1 := person {
    firstName:"Teste",
    lastName:"01",
    favoriteIceCream:[]string{"Vanilla", "Chocolate"},
  }

  p2 := person {
    firstName:"Teste",
    lastName:"02",
    favoriteIceCream:[]string{"Lemon", "Pistaccio"},
  }

  fmt.Println(p1.firstName, p1.lastName)
  for k, v := range p1.favoriteIceCream {
    fmt.Println(k, v)
  }
  fmt.Println(p2.firstName, p2.lastName)
  for k, v := range p2.favoriteIceCream {
    fmt.Println(k, v)
  }

}
