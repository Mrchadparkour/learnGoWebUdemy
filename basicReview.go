package main

import "fmt"

//package level scope
//var keyword sets to zero value
var y int

//structs -> create custom types
//type name struct { fieldNames fieldData }
//can variables with these types
//var p person
//p = data
type person struct {
  //casing controls scope
  //outside package
  Fname string
  //local
  lname string
}

//method
//p = this variable
//person = type
//speak() method name
func (p person) speak() {
  fmt.Println(p.Fname, `says, "Good Morning chickens"`)
}


type secretAgent struct {
  person
  licenseToKill bool
}

func (sa secretAgent) swag() {
  fmt.Println(sa.Fname, sa.lname, `says, "Shaken, not stirred."`)
}

//defines behavior
//both secretAgent and person have speak
//iimplies they are of similar type
//polymorphic AF
type human interface {
  speak()
}

//can call methods with its interface method
func saySomething(h human) {
  h.speak();
}

func main() {
  //short var dec operator
  //must be used in code block
  //local scoping
  //use whenever you can
  x := 7
  fmt.Println(x, y)

  //composite literal = type { data }
  xi := []int{2, 4, 7, 8, 9 }
  fmt.Println(xi)

  //maps = list of key value pairs -> map[type]valType { data }
  //reqs trailing comma
  m := map[string]int{
    "chad": 21,
    "kI": 20,
  }
  fmt.Println(m)

  p1 := person { "Chadwick", "Platt-kuhn" }

  fmt.Println(p1)

  p1.speak();

  sa1 := secretAgent{
    person {
      "James",
      "Bond",
    },
    true,
  }

  sa1.swag();
  sa1.person.speak();

  saySomething(sa1)
  saySomething(p1)

}
