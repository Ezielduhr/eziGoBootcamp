package main

import (
	stuff "example/project/first"
	"fmt"
	"reflect"
)

var pl = fmt.Println

func main() {
	pl("Hello", stuff.Name)
	intArr := []int{2, 3, 5, 7, 11}
	strArr := stuff.IntArrToStrArr(intArr)
	pl(strArr)
	pl(reflect.TypeOf(strArr))

}
