package main

import (
	"fmt"
	"math"
	"math/rand"
)

var pl = fmt.Println

func main() {
	pl(5 + 4)
	pl(5 - 4)
	pl(5 * 4)
	pl(5 / 4)
	pl(5 % 4)

	mInt := 1
	mInt += 1
	mInt++

	pl(0.1111111 + 0.1111111)

	// seedSecs := time.Now().Unix()
	// rand.Seed(seedSecs)
	randNum := rand.Intn(50) + 1
	pl("Random", randNum)
	pl("absolute", math.Abs(-10))
	pl("Power", math.Pow(4, 2))
	pl("Square", math.Sqrt(16))
	pl("Cube root", math.Cbrt(8))
	pl("rounded up", math.Ceil(4.4))
	pl("rounded down", math.Floor(4.4))
	pl("traditional round", math.Round(4.4))
	pl("Log2", math.Log2(8))
	pl("Log10", math.Log10(100))
	pl("Straight Log", math.Log(7.389))
	pl("Maximum number", math.Max(5, 4))
	pl("Minimum number", math.Min(5, 4))

	r90 := 90 * math.Pi / 180
	d90 := r90 * (180 / math.Pi)
	pl(r90, "Radians are", d90, "degrees")

}
