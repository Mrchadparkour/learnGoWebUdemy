package main

import (
  "os"
  "log"
  "text/template"
)

var tpl *template.Template

func init() {
  tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

type hotel struct {
  Name string
  Address string
  City string
  Zip int
  Region string
}

func main(){
  hotels := []hotel {
    hotel {
      "MegaHotel",
      "somPlace Rd.",
      "Vacaville",
      95688,
      "Northern",
    },
    hotel {
      "Kias Place",
      "carpenter st",
      "Redding",
      88282,
      "Northern",
    },
    hotel {
      "el gran model",
      "Gangsta place.",
      "San Diego",
      12345,
      "Southern",
    },
  }

  err := tpl.Execute(os.Stdout, hotels)
  if err != nil {
    log.Fatalln(err)
  }
}
