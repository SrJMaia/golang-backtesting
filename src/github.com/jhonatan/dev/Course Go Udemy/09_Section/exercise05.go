package main

import (
  "fmt"
  "math"
)

type circle struct {
  r float64
}

type square struct {
  l float64
}

func (c circle) area() float64 {
  return math.Pi * math.Pow(c.r, 2)
}

func (s square) area() float64 {
  return s.l * s.l
}

type shape interface {
  area() float64 // Interface needs the return type
}

func info(s shape) {
  fmt.Println(s.area())
  }

func main() {
  // Exercise 05
  // • create a type SQUARE
  // • create a type CIRCLE
  // • attach a method to each that calculates AREA and returns it
  //     ◦ circle area= π r 2
  //     ◦ square area = L * W
  // • create a type SHAPE that defines an interface as anything that has the AREA method
  // • create a func INFO which takes type shape and then prints the area
  // • create a value of type square
  // • create a value of type circle
  // • use func info to print the area of square
  // • use func info to print the area of circle

  c1 := circle{
    r:1.4142,
  }

  s1 := square{
    l:5.12,
  }

  info(c1)
  info(s1)


}
