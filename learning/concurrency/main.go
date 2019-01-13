package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("concurrency main")
	wg.Add(4)
	go startWebServer()
	go startMessageServer()
	go startPrintServer()
	go startDocumentServer()

	process()

	wg.Wait()

	moreProcessing()

}

func startWebServer() {
	time.Sleep(time.Minute)
	fmt.Println("started web server")
	wg.Done()
}

func startMessageServer() {
	fmt.Println("started message server")
	wg.Done()
}

func startPrintServer() {
	time.Sleep(time.Minute)
	fmt.Println("started print server")
	wg.Done()
}

func startDocumentServer() {
	fmt.Println("started document server")
	wg.Done()
}

func process() {
	fmt.Println("doing this while I wait for background tasks to complete")
}

func moreProcessing() {
	fmt.Println("normal processing now")
}
