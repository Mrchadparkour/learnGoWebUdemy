package main

import (
  "strings"
  "text/template"
  "log"
  "os"
)

var tpl *template.Template

var fm = template.FuncMap{
  "uc": strings.ToUpper,
  "ft": firstThree,
}

func firstThree(s string) string {
  s = strings.TrimSpace(s)
  s = s[:3]
  return s
}

type sage struct {
  Name string
  Motto string
}

type car struct {
  Manufacturer string
  Model string
  Doors int
}

func init() {
  tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
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

  sages := []sage{ maxHenry, jimmyP, callumPowell, dylanBaker }
  cars := []car { crv, bugatti }

  data := struct {
    Peeps []sage
    Trans []car
  }{
    sages,
    cars,
  }

  err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml",data)
  if err!= nil {
    log.Fatalln(err)
  }



}
