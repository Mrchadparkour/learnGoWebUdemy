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
	sages := map[string]string {
		"florida" : "Jimmy",
		"brighton": "Callum",
		"Colorado": "Dylan",
		"Van": "Max",
	}

	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}
