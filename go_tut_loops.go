package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {
	for x := 5; x >= 1; x-- {
		pl(x)
	}

	aNums := []int{1, 2, 3}
	for index, num := range aNums {
		pl(index, num)
	}

}
