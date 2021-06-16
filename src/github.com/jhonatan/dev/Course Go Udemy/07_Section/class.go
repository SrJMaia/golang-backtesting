package main

import (
  "fmt"
)

func main() {
  // ARRAYS
  var array01 [5]int
  fmt.Println(array01)
  array01[3] = 42
  fmt.Println(array01)
  fmt.Println(len(array01))

  // SLICES - COMPOSITE LITERAL
  // slice01 := type{values} // composite literal
  slice01 := []int{10, 20, 30 , 40, 50}
  fmt.Println(slice01)
  for i, v := range slice01 {
      fmt.Println(i, v)
  }

  fmt.Println(slice01[1:])
  fmt.Println(slice01[3:])
  fmt.Println(slice01[:5])

  slice02 := append(slice01, 60, 70, 80)
  fmt.Println(slice02)

  slice03 := append(slice01, slice02...)
  fmt.Println(slice03)

  for k, v := range slice03 {
    fmt.Println(k, v)
  }

  // Deleting
  slice04 := append(slice03[:2], slice03[4:]...)
  fmt.Println(slice04)

  // Using make | Type, Len, Capacity
  // Len = The size is created
  // Capacity = If the capacity if full, it would be double
  slice05 := make([]int, 10, 100)
  fmt.Println(slice05)
  fmt.Println(len(slice05))
  fmt.Println(cap(slice05))

  // Multidimensional slice
  oneD01 := []int{10, 20, 30, 40}
  fmt.Println(oneD01)
  oneD02 := []int{10, 20, 30, 40}
  fmt.Println(oneD02)

  multi01 := [][]int{oneD01, oneD02}
  fmt.Println(multi01)

  //map its the dictionary from python
  map01 := map[string]int{
    "One":1,
    "Two":2,
    "Three":3,
    "Four":4,
    "Five":5,
  }

  fmt.Println(map01)
  fmt.Println(map01["One"])
  fmt.Println(map01["Five"])
  fmt.Println(map01["Nothing"]) // Return 0 value

  // Check if the key exists
  v, ok := map01["Nothing"]
  fmt.Println(v)
  fmt.Println(ok)

  if v, ok := map01["Something"]; ok {
    fmt.Println(v)
  } else if v, ok := map01["One"]; ok {
    fmt.Println(v)
  }

  for k, v := range map01 {
    fmt.Println(k, v)
  }

  delete(map01, "One")
  fmt.Println(map01)
  delete(map01, "One Hundred") // You can delete something that does not exist
  fmt.Println(map01)

  if v, ok := m["One Hundred"]; ok { // Checks to delete
    fmt.Println("Value", v)
    delete(m, "One Hundred")
  }

}
