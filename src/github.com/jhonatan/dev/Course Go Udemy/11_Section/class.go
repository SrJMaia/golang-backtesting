package main

import (
  "encoding/json"
  "fmt"
)

type person struct {
  First string `json:"First"`// When using json must be capital letter
  Last string `json:"Last"`
  Age int `json:"Age"`
}

func main() {
  var p1 = person{
    First:"James",
    Last:"Bond",
    Age:32,
  }

  var p2 = person{
    First:"Miss",
    Last:"Moneypenny",
    Age:27,
  }

  var people = []person{
    p1,
    p2,
  }

  fmt.Println(people)

  bs, err := json.Marshal(people)
  if err != nil {
    fmt.Println(err)
  }

  fmt.Println(string(bs))

  sb := []byte(string(bs))
  fmt.Println(sb)
  fmt.Printf("%T\n", sb)

  var peopleUn = []person{}

  err = json.Unmarshal(sb, &peopleUn)
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(peopleUn)
  for k, v := range peopleUn {
    fmt.Println(k, v)
  }

}
