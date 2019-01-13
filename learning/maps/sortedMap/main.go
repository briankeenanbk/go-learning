package main

import (
	"fmt"
	"sort"
)

func main() {
	m := map[int]string{
		64: "brian",
		56: "fatiha",
		27: "john",
		59: "brendan",
	}

	fmt.Println("initialized m", m)

	/***
	  	var m map[int]string
	  	m = make(map[int] string)

	  // this would do the same
	  //m := map[int]string{}

	  	m[64] = "brian"
	  	m[56] = "fatiha"
	  	m[27] = "john"
	      m[70] = "brendan"
	  ***/

	var keys []int
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	fmt.Println("sorted by key")
	for _, k := range keys {
		fmt.Println("Key:", k, "Value:", m[k])
	}

	ms := map[string]int{}

	for k, v := range m {
		ms[v] = k
	}

	var skeys []string
	for _, v := range m {
		skeys = append(skeys, v)
	}

	sort.Strings(skeys)
	fmt.Println("sorted by value")
	for _, k := range skeys {
		fmt.Println("--Value:", k, "Key:", ms[k])
	}
}
