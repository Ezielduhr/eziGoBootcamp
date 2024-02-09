package main

import (
	"fmt"
	"log"
	"time"
)

var pl = fmt.Println

func main() {
	now := time.Now()
	pl(now.Year(), now.Month(), now.Day())
	pl(now.Local())
	location, err := time.LoadLocation("America/New_York")
	if err != nil {
		log.Fatal(err)
	}
	pl(now.In(location))

}
