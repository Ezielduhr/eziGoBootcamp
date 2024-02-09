package main

import "fmt"

var pl = fmt.Println

func main() {
	slice1 := make([]string, 6)
	slice1[0] = "Society"
	slice1[1] = "of"
	slice1[2] = "the"
	slice1[3] = "simulated"
	slice1[4] = "universe"
	slice1[5] = "party"

	pl("Slice size", len(slice1))
	for i := 0; i < len(slice1); i++ {
		pl(slice1[i])
	}

	for _, x := range slice1 {
		pl(x)
	}

	pl(slice1)

	slice2 := slice1[0:2]
	pl(slice2)
}
