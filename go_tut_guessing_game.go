package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var pl = fmt.Println

func main() {
	reader := bufio.NewReader(os.Stdin)
	pl("enter number 1")
	iNumber1, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	pl("enter number 2")
	iNumber2, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	iNumber1 = strings.TrimSpace(iNumber1)
	iNumber2 = strings.TrimSpace(iNumber2)
	number1, err := strconv.Atoi(iNumber1)
	number2, err := strconv.Atoi(iNumber2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s + %s = %d\n", iNumber1, iNumber2, (number1 + number2))

}
