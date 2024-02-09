package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

var pl = fmt.Println

func main() {
	sV1 := "A word"
	replacer := strings.NewReplacer("A", "Another")
	sV2 := replacer.Replace(sV1)

	// Escape sequences \n \t \ \\

	pl(sV2)
	pl("Length", len(sV2))
	pl("Contains Another", strings.Contains(sV2, "Another"))
	pl("o index", strings.Index(sV2, "o"))
	pl("Replace", strings.Replace(sV2, "o", "0", -1))
	pl("Split", strings.Split("a-b-c-d", "-"))
	pl("Lower", strings.ToLower(sV2))
	pl("Has prefix", strings.HasPrefix("TacoCat", "Taco"))

	rTest := "abcdefge"
	pl("Rune count", utf8.RuneCountInString(rTest))
	for _, rValue := range rTest {
		pl(rValue)
	}

}
