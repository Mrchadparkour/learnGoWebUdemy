package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

//init is called before main
func init() {
	//go and get all gohtml files from templates folder
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	write("one.gohtml")
	write("chicken.gohtml")
	write("potato.gohtml")
}

func write(filename string) {
	path := "public/" + filename[0:len(filename)-7] + ".html"
	nf, err := os.Create(path)
	if err != nil {
		log.Fatalln(err)
	}
	defer nf.Close()

	err = tpl.ExecuteTemplate(nf, filename, 42)
	if err != nil {
		log.Fatalln(err)
	}
}
