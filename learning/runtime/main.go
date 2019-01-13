package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOOS, runtime.GOARCH)
	fmt.Println(runtime.GOROOT())
	fmt.Println(runtime.NumCPU())
}
