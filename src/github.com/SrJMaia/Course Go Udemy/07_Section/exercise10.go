package main

import (
  "fmt"
)

func main() {
  // Exercise 10
  // Using the code from the previous example, delete a record from your map.
  // Now print the map out using the “range” loop

  map01 := map[string][]string{
    "bond_james":     []string{"Shaken, not stirred","Martinis","Women"},
    "moneypenny_miss":[]string{"James Bond","Literature","Computer Science"},
    "no_dr":          []string{"Being evil","Ice cream", "Sunsets"},
  }

  map01["fleming_ian"] = []string{"steaks", "cigars", "espionage"}

  delete(map01, "bond_james")

  for k, v := range map01 {
    fmt.Println("Index: ", k)
    for k2, v2 := range v {
      fmt.Println("\tValue: ", v2, "Key: ", k2)
    }
  }

}
