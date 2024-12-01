package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	var x float64

	fmt.Print("Введите натуральное число n: ")
	fmt.Scan(&n)
	fmt.Print("Введите действительное число x: ")
	fmt.Scan(&x)

	closestValue := 0.0
	minDifference := math.MaxFloat64

	for k := 1; k <= n; k++ {
		value := math.Exp(math.Cos(x*x*float64(k))) * math.Sin(x*x*x*float64(k))
		roundedValue := math.Round(value)
		difference := math.Abs(value - roundedValue)

		if difference < minDifference {
			minDifference = difference
			closestValue = value
		}
	}

	fmt.Printf("Ближайшее к целому значение: %f\n", closestValue)
}
