package main

import (
	"fmt"
	"math/rand"
)

func shellSort(arr []int) {
    n := len(arr)
    gap := n / 2

    for gap > 0 {
        for i := gap; i < n; i++ {
            temp := arr[i]
            j := i
            for j >= gap && arr[j-gap] > temp {
                arr[j] = arr[j-gap]
                j -= gap
            }
            arr[j] = temp
        }
        gap /= 2
    }
}

func main() {
    X := make([][]int, 5)
	for i := range X {
		X[i] = make([]int, 7)
		for j := range X[i] {
			X[i][j] = rand.Intn(100)
		}
	}

	for i := 0; i < len(X); i++ {
        fmt.Println(X[i])
    }

    for i := 0; i < len(X); i++ {
        shellSort(X[i])
    }

	fmt.Println("\nSorted:")

    for i := 0; i < len(X); i++ {
        fmt.Println(X[i])
    }
}
