package main

import "fmt"

var pl = fmt.Println

func changeVal2(myPointer *int) {
	*myPointer = 12
}

func dubbleArrayValues(arr *[4]int) {
	for x := 0; x < 4; x++ {
		arr[x] *= 2
	}
}

func main() {
	f4 := 10
	pl("f4 value:", f4)
	pl("f4 address:", &f4)
	var f4Pointer *int = &f4
	pl("f4 address:", f4Pointer)
	pl("f4 value:", *f4Pointer)

	*f4Pointer = 11
	pl("f4 value:", *f4Pointer)
	pl("f4 value:", f4)

	changeVal2(f4Pointer)
	pl("f4 value:", f4)

	array1 := [4]int{2, 4, 8, 16}
	dubbleArrayValues(&array1)
	pl(array1)

}
