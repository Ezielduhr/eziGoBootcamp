package main

import (
	"fmt"
	"log"
)

var pl = fmt.Println

func sayHello() {
	pl("Hello")
}

func getSum(x int, y int) int {
	return x + y
}

func getQuotient(x float64, y float64) (ans float64, err error) {
	if y == 0 || x == 0 {
		return 0, fmt.Errorf("can't devide by zero")
	} else {
		return x / y, nil
	}
}

func getTwo(x int) (int, int) {
	return x + 1, x + 2
}

func getSumN(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func getArraySum(arr []int) int {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	return sum
}

func main() {
	sayHello()
	pl(getSum(2, 5))

	// ans, err := getQuotient(5, 0)
	ans, err := getQuotient(5, 4)
	if err != nil {
		log.Fatal(err)
	}
	pl(ans)

	pl(getTwo(2))
	pl(getSumN(2, 5, 7, 8))
	numberArray := []int{2, 5, 6, 100}
	pl(getArraySum(numberArray))
}
