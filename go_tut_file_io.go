package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

var pl = fmt.Println

func main() {
	log.SetFlags(log.Lshortfile)
	file, err := os.Create("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	iPrimerArr := []int{2, 3, 5, 7, 11}
	var sPrimeArr []string
	for _, i := range iPrimerArr {
		sPrimeArr = append(sPrimeArr, strconv.Itoa(i))
	}
	for _, num := range sPrimeArr {
		_, err := file.WriteString(num + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
	file, err = os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	scan1 := bufio.NewScanner(file)
	for scan1.Scan() {
		pl("Prime :", scan1.Text())
	}
	if err := scan1.Err(); err != nil {
		log.Fatal(err)
	}
	// os.Remove("data.txt")
	_, err = os.Stat("data.txt")
	if errors.Is(err, os.ErrNotExist) {
		log.Fatal(err)
	}

	file, err = os.OpenFile("data.txt", os.O_APPEND|os.O_CREATE|os.O_RDONLY, 0644)

}
