package main
import "fmt"
import "github.com/aminmeyghani/app1/mlib/math"
import "github.com/aminmeyghani/app1/mlib/simple"
import "github.com/aminmeyghani/app1/mlib/closure"
import "github.com/aminmeyghani/app1/mlib/mol"
import "github.com/aminmeyghani/app1/mlib/objects"
import "github.com/aminmeyghani/app1/mlib/basics"
import "container/list"

func main() {
  // test the average function
  xs := []float64{1,2,3,4}
  avg := math.Average(xs)
  fmt.Println(avg)
  // test the simple function that just returns a single float value.
  fmt.Println(simple.Simple())
  // test the closure application
  makeInc := closure.IncMaker()
  fmt.Println(makeInc())
  fmt.Println(makeInc())
  fmt.Println(makeInc())
  // test the list
  fmt.Println("............ linkedlist")
  // Create a new list and put some numbers in it.
  l := list.New()
  e4 := l.PushBack(4)
  e1 := l.PushFront(1)
  l.InsertBefore(3, e4)
  l.InsertAfter(2, e1)

  // Iterate through list and print its contents.
  for e := l.Front(); e != nil; e = e.Next() {
    fmt.Println(e.Value)
  }

  // does avo exist
  fmt.Println(mol.MyConstant)

  // moles and atoms
  mol.PrintAll()

  // print the list of objects
  objects.PrintCars()

  // print the list of users
  objects.PrintUsers()

  // the basics
  fmt.Println("BASICS______________________")
  basics.Eg1()
  basics.Eg2()
}
