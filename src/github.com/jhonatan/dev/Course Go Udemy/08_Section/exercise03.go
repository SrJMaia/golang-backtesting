package main

import (
  "fmt"
)

type new_vehicle struct {
  color string
  doors uint8
}

type truck struct {
  new_vehicle
  fourWheel bool
}

type sedan struct {
  new_vehicle
  luxury bool
}

func main() {
  // Exercise 03
  // • Create a new type: vehicle.
  //     ◦ The underlying type is a struct.
  //     ◦ The fields:
  //         ▪ doors
  //         ▪ color
  // • Create two new types: truck & sedan.
  //     ◦ The underlying type of each of these new types is a struct.
  //     ◦ Embed the “vehicle” type in both truck & sedan.
  //     ◦ Give truck the field “fourWheel” which will be set to bool.
  //     ◦ Give sedan the field “luxury” which will be set to bool. solution
  // • Using the vehicle, truck, and sedan structs:
  //     ◦ using a composite literal, create a value of type truck and assign values to the fields;
  //     ◦ using a composite literal, create a value of type sedan and assign values to the fields.
  // • Print out each of these values.
  // • Print out a single field from each of these values.

  truck01 := truck {
    new_vehicle:new_vehicle {
      color:"Red",
      doors:2,
    },
    fourWheel:false,
  }

  sedan01 := sedan {
    new_vehicle:new_vehicle {
      color:"Black",
      doors:4,
    },
    luxury:true,
  }

  fmt.Println(truck01.color, truck01.doors, truck01.fourWheel)
  fmt.Println(sedan01.color, sedan01.doors, sedan01.luxury)
}
