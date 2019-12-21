package main

import (
	"os"
	"text/template"
)

func main() {
	t, err := template.ParseFiles("learning/web/templates/01/layout.tmpl")
	if err != nil {
		panic(err)
	}

	nf, err := os.Create("learning/web/templates/01/output.txt")
	defer nf.Close()

	nf.WriteString("below from the template\n")

	err = t.Execute(nf, "hello from GO again")
	if err != nil {
		panic(err)
	}

	nf.WriteString("\nThis isn't from the template")
}
