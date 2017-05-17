// exploring the basics of the language.
package basics
import "fmt"

func Eg1() {
  // type inference, basic assignments
  var name = "Tom"
  var triple = [3]int{1,2,3}
  elms := []string{"tom", "bob", "jon"}
  isMetal := false
  fmt.Println(name, triple, elms, isMetal)
}

func Eg2() {
  // composite types: array, slice, map
  // slice: variable sizs array
  type toy struct {
    id int
    name string
  }
  var toys = []toy {
    toy{5, "a"},
    toy{15, "b"},
  }

  // make a struct and initialize it
  planet := struct {
    name string
    diameter int
  }{
    "earth", 12773,
  }

  fmt.Println(toys, planet)
}

func Obj() {
  // method is a function attached to a custom type.
  type amu float64
  func (mass amu) float() float64 {
    return float64(mass)
  }
}


func InterfaceExample() {
  type Stringer interface {
    String() string
  }
  // any type that has the method String() attached
  // automatically implements the Stringer interface.
  type metal struct {
    name string
    number int32
    weight amu
  }
  func (m metalloid) String() string {
    // ...
  }
}





