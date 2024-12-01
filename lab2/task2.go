package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	X := make([][]int, 8)
	for i := range X {
		X[i] = make([]int, 10)
		for j := range X[i] {
			X[i][j] = rand.Intn(100)
		}
	}

	for _, row := range X {
		fmt.Println(row)
	}

	for j := 0; j < len(X[0]); j++ {
		for i := 1; i < len(X); i++ {
			key := X[i][j]
			k := i - 1
			for k >= 0 && X[k][j] > key {
				X[k+1][j] = X[k][j]
				k = k - 1
			}
			X[k+1][j] = key
		}
	}

	lastRow := X[len(X)-1]
	for i := 1; i < len(lastRow); i++ {
		key := lastRow[i]
		j := i - 1
		for j >= 0 && lastRow[j] > key {
			lastRow[j+1] = lastRow[j]
			j = j - 1
		}
		lastRow[j+1] = key
	}

	fmt.Println("\nSorted:")

	for _, row := range X {
		fmt.Println(row)
	}
}
