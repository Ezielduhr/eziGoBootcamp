package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

var pl = fmt.Println

func main() {
	pl(os.Args)
	args := os.Args[1:]

	var iArgs = []int{}
	for _, i := range args {
		val, err := strconv.Atoi(i)
		if err != nil {
			log.Fatal(err)
		} else {
			iArgs = append(iArgs, val)
		}
	}
	maxVal := 0
	for _, val := range iArgs {
		if val > maxVal {
			maxVal = val
		}
	}
	pl("Max: ", maxVal)

}

/*
go build
./go_tut_cli
cp go_tut_cli $GOPATH/bin/go_tut_cli
*/
