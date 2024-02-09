package main

import (
	"fmt"
	"reflect"
)

var pl = fmt.Println

type TeaSpoon float64
type TableSpoon float64
type MiliLiters float64

func TeaSpoonToMiliLiters(tsp TeaSpoon) MiliLiters {
	return MiliLiters(tsp * 4.92)
}

func TableSpoonToMiliLiters(tbs TableSpoon) MiliLiters {
	return MiliLiters(tbs * 14.79)
}

func (tsp TeaSpoon) ToMls() MiliLiters {
	return MiliLiters(tsp * 4.92)
}

func (tbs TableSpoon) ToMls() MiliLiters {
	return MiliLiters(tbs * 14.79)
}

func main() {
	// bad way
	conversion1 := TeaSpoonToMiliLiters(3)
	pl(conversion1)
	pl(reflect.TypeOf(conversion1))

	// better way
	conversion2 := TeaSpoon(3).ToMls()
	pl(conversion2)
	pl("^ Teaspoon to MLs")

	conversion3 := TableSpoon(3).ToMls()
	pl(conversion3)
	pl("^ TableSpoon to MLs")

}
