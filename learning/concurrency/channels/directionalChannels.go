package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	go send(67, c)
	//go send(69, c)
	//go send(70, c)

	i := receive(c)

	fmt.Println(i)

}

func send(i int, c chan<- int) {
	time.Sleep(time.Minute)
	fmt.Println("sending", c)
	c <- i
	i++
	c <- i
	i++
	c <- i
	close(c)
}

func receive(c <-chan int) int {
	var i int
	for v := range c {
		i = v
		fmt.Println("recieving range", i)
	}

	return i
}
