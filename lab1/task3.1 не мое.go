package main

import "fmt"

func denominator(i, j float64) float64 {
	if j >= 22.0 {
		return 0
	}
	return j + i/denominator(i, j + 4.0)
}

func calculate(x, y float64) float64 {
	return (y + x) / (y - x)
}

func main() {
	for i := 0.0; i <= 1.0; i += 0.01 {
		y := denominator(i, 2.0)
		f := calculate(i, y)

		fmt.Printf("y: %f\n", y)
		fmt.Printf("x: %f\n", i)
		fmt.Printf("f: %f\n", f)
	}
}
