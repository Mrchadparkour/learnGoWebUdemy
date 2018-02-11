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

type car struct {
  Manufacturer string
  Model string
  Doors int
}

// type items struct {
//   Wisdom []sage
//   Transport []car
// }

func init() {
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

  crv := car {
    Manufacturer: "Honda",
    Model: "crv",
    Doors: 4,
  }

  bugatti := car {
    Manufacturer: "bugatti",
    Model: "Boogee AF",
    Doors: 2,
  }

  sages := []sage {maxHenry, jimmyP, callumPowell, dylanBaker}
  cars := []car {crv, bugatti}

  //Singlue use struct style
  data := struct {
    Wisdom []sage
    Transport []car
  }{
    sages,
    cars,
  }

  err := tpl.Execute(os.Stdout, data)
  if err != nil {
    log.Fatalln(err)
  }

}
