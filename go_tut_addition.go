package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

var pl = fmt.Println

func main() {
	numberRange := 50
	randNumber := rand.Intn(numberRange) + 1
	reader := bufio.NewReader(os.Stdin)
	guessing := true

	for guessing {
		fmt.Printf("Guess a number in the range 0 %d:\n", numberRange)
		guessedNumber, err := reader.ReadString('\n')
		guessedNumber = strings.TrimSpace(guessedNumber)
		iGuessedNumber, err := strconv.Atoi(guessedNumber)
		if err != nil {
			log.Fatal(err)
		}

		if iGuessedNumber > randNumber {
			pl("Guess lower")
		} else if iGuessedNumber < randNumber {
			pl("Guess higher")
		} else if iGuessedNumber == randNumber {
			pl(randNumber, "is the number, you've won ")
			guessing = false
		} else {
			pl("what are you doing?")
		}

	}
}
