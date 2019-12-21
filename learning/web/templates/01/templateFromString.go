package main

import (
	"html/template"
	"os"
)

func main() {
	t, err := template.New("test").Parse("This is what I passed you {{.}}")
	if err != nil {
		panic(err)
	}

	//err = t.Execute(os.Stdout, 56)
	//if err != nil {
	//	panic(err)
	//}

	t, err = t.Parse("Now parse this {{.}}")
	if err != nil {
		panic(err)
	}

	err = t.Execute(os.Stdout, "Ah, Im a string")
	if err != nil {
		panic(err)
	}

}
