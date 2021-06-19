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
  // Exercise 02
  // Take the code from the previous exercise, then store the values of
  // type person in a map with the key of last name.
  // Access each value in the map. Print out the values, ranging over the slice.

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

  map01 := map[string]person{
    p1.lastName: p1,
    p2.lastName: p2,
  }

  for _, v := range map01 {
    fmt.Println(v.firstName)
    fmt.Println(v.lastName)
    for kk, vv := range v.favoriteIceCream {
      fmt.Println(kk, vv)
    }
    fmt.Println("------------")
  }

}
