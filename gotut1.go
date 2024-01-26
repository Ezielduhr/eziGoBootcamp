package main

import (
	"bufio"
	f "fmt"
	"log"
	"os"
)

var pl = f.Println

// Comment
/*
block comment
*/

func main() {
	pl("Hello Go")
	pl("What is your name?")

	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')

	if err == nil {
		pl("Hello", name)
	} else {
		log.Fatal(err)
	}

}
