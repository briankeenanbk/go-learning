package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	t, err := template.ParseFiles("learning/web/templates/01/layout.tmpl", "learning/web/templates/01/layout2.tmpl")
	if err != nil {
		panic(err)
	}

	err = t.ExecuteTemplate(os.Stdout, "layout.tmpl", "Hello")
	if err != nil {
		log.Fatalln(err)
	}

	err = t.ExecuteTemplate(os.Stdout, "layout2.tmpl", "Hello there")
	if err != nil {
		log.Fatalln(err)
	}
}
