package main

import (
	"fmt"
	"log"
)

type person struct {
	first string
	last  string
}

func (n person) Error() string {
	return fmt.Sprintf("a norgate math error occured: %v %v", n.first, n.last)
}

func main() {
	p := person{"brian", "keenan"}
	_, err := sqrt(-10.23)
	fmt.Println(p)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		nme := fmt.Errorf("norgate math redux: square root of negative number: %v", f)
		return 0, nme
	}
	return 42, nil
}

// see use of structs with error type in standard library:
//
// http://golang.org/pkg/net/#OpError
// http://golang.org/src/pkg/net/dial.go
// http://golang.org/src/pkg/net/net.go
//
// http://golang.org/src/pkg/encoding/json/decode.go
