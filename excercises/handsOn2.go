package main

import (
  "fmt"
)

func main() {
  mike := person{"Mike", "D"}
  a313 := secretAgent{
      person{"Marshal", "Mathers"},
      313,
    }
  fmt.Println(mike.fname)

  fmt.Println(a313.codeName)

  a313.saSpeak()
  mike.pSpeak()
}



type person struct {
  fname string
  lname string
}

func (p person) pSpeak() {
  fmt.Println(p.fname, `says, "Whaddup G"`)
}

type secretAgent struct {
  person
  codeName int
}

func (sa secretAgent) saSpeak() {
  fmt.Println(sa.codeName, `says, "danger is my middle name."`)
}
