package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	//tpl is a pointer && think about it like container
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	//print to console
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	//dump to file
	nf, err := os.Create("index.html")
	if err != nil {
		log.Println("error creating file", err)
	}
	//need to learn mo0re about why defer is important
	defer nf.Close()

	err = tpl.Execute(nf, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
