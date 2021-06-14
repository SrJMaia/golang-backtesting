package main

import (
  "fmt"
)

func main() {
  // FOR LOOP
  // for init; condition; post {
  // }
  // for i := 0; i <= 10; i++ {
  //   fmt.Println(i)
  // }

  // for i := 0; i <= 5; i++ {
  //   fmt.Println(i)
  //   for j := 0; j <= 2; j++ {
  //     fmt.Println(j)
  //   }
  // }

  // WHILE BASICALLY
  // x := 1
  // for x < 10 {
  //   fmt.Println(x)
  //   x++
  // }

  // for i := 33; i < 126; i++ {
  //   fmt.Printf("%c\t%#U\t%#x\n", i, i, i)
  //   }

  // IF ELSE CONDITIONAL
  if true {
    fmt.Println("001")
  }
  if false {
    fmt.Println("002")
  }
  if !true {
    fmt.Println("003")
  }
  if !false {
    fmt.Println("004")
  }

  if x := 42; x==42 {
    fmt.Println("005")
  }
  fmt.Println("Same"); fmt.Println("Line")


  variable := 10

  if variable == 1 {
    fmt.Println("One")
  } else if variable > 1 {
    fmt.Println("Greater")
  } else {
    fmt.Println("Less")
  }

  switch {
	case false:
		fmt.Println("this should not print")
	case (2 == 4):
		fmt.Println("this should not print2")
	case (3 == 3):
		fmt.Println("prints")
		fallthrough
	case (4 == 4):
		fmt.Println("also true, does it print?")
		fallthrough
	case (7 == 9):
		fmt.Println("not true1")
		fallthrough
	case (11 == 14):
		fmt.Println("not true 2")
		fallthrough
	case (15 == 15):
		fmt.Println("true 15")
		fallthrough
	default:
		fmt.Println("this is default")
	}

  switch {
	case false:
		fmt.Println("this should not print")
	case (2 == 4):
		fmt.Println("this should not print2")
	default:
		fmt.Println("this is default")
	}

  // CONDITIONAL OPERATORS
  // && and
  // || or
  // ! not
}
