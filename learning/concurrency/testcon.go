package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	a := 7
	wg.Add(1)
	go foo(&a)
	for i := 1; i < 100; i++ {
		fmt.Println(i)
	}
	wg.Wait()
	fmt.Println(a)
}

func foo(a *int) int {
	*a += 2
	fmt.Println("ahh")
	wg.Done()
	return 67

}
