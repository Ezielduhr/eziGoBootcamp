package main

import "fmt"

var pl = fmt.Println

func main() {
	var array1 [5]int
	array1[0] = 1
	array2 := [5]int{1, 2, 3, 4, 5}
	pl("Index 0:", array2[0])
	pl("Array Length:", len(array2))

	for i := 0; i < len(array2); i++ {
		pl(array2[i])
	}

	array3 := [2][2]int{
		{1, 2},
		{3, 4},
	}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			pl(array3[i][j])
		}
	}

}
