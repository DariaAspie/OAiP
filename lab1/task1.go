package main

import (
	"fmt"
	"math"
)

func calculate(x float64, n int) float64 {
	sum := 0.0
	for k := 1; k <= n; k++ {
		term := (float64(k-1) / float64(k)) + math.Pow(math.E, math.Cos(x))
		sum += term
	}
	return math.Cbrt(float64(n)*x - sum)
}

func printResults(xStart, xEnd, deltaX float64, nStart, nEnd int) {
	for n := nStart; n <= nEnd; n++ {
		for x := xStart; x <= xEnd+0.0001; x += deltaX {
			result := calculate(x, n)
			fmt.Printf("n = %d\t", n)
			fmt.Printf("\tx = %.2f\t", x)
			fmt.Printf("f = %.6f\t\n", result)
		}
	}
}

func main() {
	xStart := 0.6
	xEnd := 1.1
	deltaX := 0.25
	nStart := 10
	nEnd := 15

	printResults(xStart, xEnd, deltaX, nStart, nEnd)
}
