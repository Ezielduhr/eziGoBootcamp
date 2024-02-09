package main

import "fmt"

var pl = fmt.Println

func main() {
	aString1 := "abcde"
	rArray := []rune(aString1)
	for _, v := range rArray {
		pl("Rune array :", v)
	}
	byteArray := []byte{'a', 'b', 'c'}
	byteString := string(byteArray[:])
	pl("I'm a string:", byteString)

}
