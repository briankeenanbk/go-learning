package main

import "fmt"

func main() {
	a := []int{2, 5, 7, 9}
	fmt.Println(a)

	for i, v := range a {
		fmt.Println(i, v)
	}

	a = append(a, 67, 89)
	fmt.Println(a)

	for i, v := range a {
		fmt.Println(i, v)
	}

	b := append(a[1:3], a[4:]...)
	for i, v := range b {
		fmt.Println(i, v)
	}

	c := make([]int, 10, 20)
	fmt.Println(c)

}
