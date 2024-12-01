package main

import "fmt"

func main() {
	for x := 0.5; x < 0.85; x += 0.05 {
		fmt.Printf("[i,j]: ")
		y := denominator(15.0, x, 0.0)

		fmt.Printf("x: %.2f\t", x)
		fmt.Printf("y: %f\t\n\n", y)
	}
}

func denominator(i, x, j float64) float64 {
	if i < 0.0 {
		fmt.Printf("\n")
		return 0
	}
	fmt.Printf("[%d, %.2f] ", int(i), j)
	return j + (factorial(i)*(x-i))/denominator(i-1.0, x, ensure_not_zero(j)*2.0)
}

func factorial(n float64) float64 {
    if n == 0.0 {
        return 1
    }
    return n * factorial(n-1.0)
}

func ensure_not_zero(n float64) float64 {
	if n == 0.0 {
		return 0.25 / 2.0
	}
	return n
}
