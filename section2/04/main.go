package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	//template.Must error checks for you
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	sages := []string{"Jimmy Perrira", "Callum Powell", "Dylan Baker", "Max Henry"}
	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}
