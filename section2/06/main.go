package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name string
	Motto string
}

func init() {
	//template.Must error checks for you
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}


func main() {
	maxHenry := sage {
		Name: "Max Henry",
		Motto: "Sexual frustration is the key to parkour",
	}

	dylanBaker := sage {
		Name: "Dylan Baker",
		Motto: "Man hood was pretty scary",
	}

	callumPowell := sage {
		Name: "Callum Powell",
		Motto: "Lift More",
	}

	jimmyP := sage {
		Name: "Jimmy Perrira",
		Motto: "How do worms dig so good?",
	}

	sages := []sage {maxHenry, dylanBaker, callumPowell, jimmyP}

	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}
