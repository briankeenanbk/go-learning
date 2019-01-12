package main

import "fmt"

type speakable interface {
	talk()
}

type person struct {
	name   string
	saying string
}

func (p person) talk() {
	fmt.Println(p.name, "says", p.saying)
}

func talk(p speakable) {
	p.talk()
}

func main() {
	p := person{"brian", "hello"}
	p.talk()
	talk(p)
}
