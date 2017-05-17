// make a list of structs and print them
package objects
import "fmt"

// define a car type
type car struct {
  color string
  name string
  doors int
}

// make a list of cars
var cars = []car {
  car{"blue", "tom", 4},
  car{"red", "jon", 5},
}

// define a user type
type user struct {
  id int
  name string
}

var users = []user {
  user{5, "tom"},
  user{11, "jonny"},
}


func PrintCars() {
  for _, m := range cars{
    fmt.Println(m.name)
    fmt.Println(m.color)
    fmt.Println(m.doors)
  }
}

func PrintUsers() {
  fmt.Println("printing users ----------------")
  for _, u := range users {
    fmt.Println(u.id)
    fmt.Println(u.name)
  }
}
