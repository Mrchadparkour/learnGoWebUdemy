package main

import (
	"fmt"
)

func main() {
	sa := secretAgent{
		person{
			"Marshal",
			"Mathers",
		},
		313,
	}
	p1 := person{"Cheddar", "Bob"}

	spitFire(p1)
	spitFire(sa)
}

type human interface {
	rap()
}

func spitFire(h human) {
	h.rap()
}

type person struct {
	fname string
	lname string
}

func (p person) pSpeak() {
	fmt.Println(p.fname, `says, "Whaddup G"`)
}

func (p person) rap() {
	fmt.Println(p.fname, "Hippity hoppity.")
}

type secretAgent struct {
	person
	codeName int
}

func (sa secretAgent) saSpeak() {
	fmt.Println(sa.codeName, `says, "danger is my middle name."`)
}

func (sa secretAgent) rap() {
	fmt.Println(sa.fname, sa.codeName, "Mom's spaghetti is on my sweater already.")
}
