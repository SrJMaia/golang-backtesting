package main

import (
  "fmt"
)

func main() {
  // Exercise 07
  // Create a slice of a slice of string ([][]string).
  // Store the following data in the multi-dimensional slice:
  //
  // "James", "Bond", "Shaken, not stirred"
  // "Miss", "Moneypenny", "Helloooooo, James."
  //
  // Range over the records, then range over the data in each record.

  one := []string{"James", "Bond", "Shaken, not stirred"}
  two := []string{"Miss", "Moneypenny", "Helloooooo, James."}

  multi := [][]string{one, two}

  for k, v := range multi {
    fmt.Println(k)
    for _, v2 := range v {
      fmt.Println("\t", v2)
    }
  }

}
