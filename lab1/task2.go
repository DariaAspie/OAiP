package main

import (
	"fmt"
	"math"
)

func main() {
	for i := 1000; i <= 9999; i++ {
		firstTwo := i / 100
		lastTwo := i % 100
		sqrtLastTwo := math.Sqrt(float64(lastTwo))

		if math.Sqrt(float64(i)) == float64(firstTwo)+sqrtLastTwo {
			fmt.Printf("Found number: %d\n", i)
		}
	}
}
