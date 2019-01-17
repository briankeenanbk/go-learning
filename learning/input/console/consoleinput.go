package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)

	fmt.Println("Scan Enter text: ")
	var text2 string
	fmt.Scanln(&text2)
	fmt.Println(">", text2, "<")

	ln := text2
	fmt.Sscanln("%v", ln)
	fmt.Println(ln)

	fmt.Println("=======================================")

	readIt()
	scanner()
}

func readIt() {
	fmt.Println("Readint It")
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()

	if err != nil {
		fmt.Println(err)
	}

	// print out the unicode value i.e. A -> 65, a -> 97
	fmt.Println(char)

	switch char {
	case 'A':
		fmt.Println("A Key Pressed")
		break
	case 'a':
		fmt.Println("a Key Pressed")
		break
	}
}

func scanner() {
	fmt.Println("====Scanner===")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
