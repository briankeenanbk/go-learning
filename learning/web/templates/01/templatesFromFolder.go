package main

import (
	"os"
	"text/template"
)

func main() {
	t, err := template.ParseGlob("learning/web/templates/01/*.tmpl")
	if err != nil {
		panic(err)
	}

	t.Execute(os.Stdout, []string{"hello", "there"})

	t.Execute(os.Stdout, "just this")

	//err = t.ExecuteTemplate(os.Stdout, "layout2.tmpl", "hi")
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = t.ExecuteTemplate(os.Stdout, "layout.tmpl", "hi there")
	//if err != nil {
	//	panic(err)
	//}

}
