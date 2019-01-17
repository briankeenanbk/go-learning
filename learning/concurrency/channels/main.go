package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

func main() {
	p1 := person{"brian", "keenan"}
	p2 := person{"fatiha", "jamil"}
	ca := make(chan person, 2)
	ca <- p1
	ca <- p2
	fmt.Println(<-ca)
	fmt.Println(<-ca)

	c := make(chan int)
	cs := make(chan<- int)

	go func() {
		cs <- 42
	}()
	cr := make(<-chan int) // receive

	fmt.Println("-----")
	fmt.Printf("c\t%T\n", c)
	fmt.Printf("cr\t%T\n", cr)
	fmt.Printf("cs\t%T\n", cs)

	cr = c

	fmt.Println("-----")
	fmt.Printf("c\t%T\n", c)
	fmt.Printf("cr\t%T\n", cr)
	fmt.Printf("cs\t%T\n", cs)

	//fmt.Println(<-cs)
}
