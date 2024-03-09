package main

import "fmt"

var pl = fmt.Println

func factorial(num uint64) uint64 {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}


func main() {
	// Recursion
	pl("Factorial 4 =", factorial(4))
}