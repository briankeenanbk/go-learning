package main

import (
	"fmt"
	maybe "github.com/briankeenanbk/go-learn/learning"
)

//main
func main() {
	var value maybe.Export
	value.DoMagic()
	fmt.Printf("%s", value)

}
