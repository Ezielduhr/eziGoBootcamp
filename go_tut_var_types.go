package main

import (
	"fmt"
	"reflect"
	"strconv"
)

var pl = fmt.Println

func main() {
	/*
		int, float64, bool, string, rune
		Default type 0, 0.0, false, ""
		= assigning new value to variable
		:= creates a new variable

	*/
	pl(reflect.TypeOf(25))
	pl(reflect.TypeOf(3.14))
	pl(reflect.TypeOf(true))
	pl(reflect.TypeOf("Hello"))

	cV1 := 1.5
	cV2 := int(cV1)
	pl(cV2)

	cV3 := "500000"
	cV4, err := strconv.Atoi(cV3)
	pl(cV4, err, reflect.TypeOf(cV4))

	cV5 := 5000000
	cV6 := strconv.Itoa(cV5)
	pl(cV6)

	cV7 := "3.14"
	if cV8, err := strconv.ParseFloat(cV7, 64); err == nil {
		pl(cV8)
	}

	cV9 := fmt.Sprintf("%f", 3.14)
	pl(cV9)

}
