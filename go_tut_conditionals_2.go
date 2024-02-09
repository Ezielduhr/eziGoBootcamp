package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("What is your age?")
	reader := bufio.NewReader(os.Stdin)
	iAge, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	iAgeString := strings.ReplaceAll(iAge, " ", "")
	iAgeString = strings.TrimSuffix(iAgeString, "\n")
	iAgeClean, err := strconv.Atoi(iAgeString)
	if err != nil {
		log.Fatal(err)
	}

	if iAgeClean < 5 {
		fmt.Println("Too young for school")
	} else if iAgeClean == 5 {
		fmt.Println("Go to Kindergarten")
	} else if (iAgeClean > 5) && (iAgeClean <= 17) {
		grade := iAgeClean - 5
		fmt.Println("Go to grade", grade)
	} else {
		fmt.Println("Go to college")
	}
}
