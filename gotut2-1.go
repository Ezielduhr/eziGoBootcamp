package main

import f "fmt"

var pl = f.Println

func main() {
	/*
		If Conditional
		Conditional operators:  > < >= <= == !=
		Logical operators: && || !
	*/

	iAge := 8
	if (iAge >= 1) && (iAge <= 18) {
		pl("Important Birthday")
	} else if (iAge == 21) || (iAge == 50) {
		pl("Important Birthday")
	} else if iAge >= 65 {
		pl("Important Birthday")
	} else {
		pl("Not an important Birthday")
	}
}
