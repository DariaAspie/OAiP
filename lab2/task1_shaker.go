package main

import (
	"fmt"
	"math/rand"
	"time"
	"sort"
)

func main() {
	random_arr := generateRandomArray()

	// методы модуля sort переопределяют переменную, поэтому приходится копировать массив
	sorted_arr := make([]int, len(random_arr))
	copy(sorted_arr, random_arr)
	sort.Ints(sorted_arr)

	inverted_arr := make([]int, len(sorted_arr))
	copy(inverted_arr, sorted_arr)
	sort.Sort(sort.Reverse(sort.IntSlice(inverted_arr)))

	sliceSizes := []int{100, 200, 300, 500, 1000, 2000}

	fmt.Printf("Random:\n")
	for _, size := range sliceSizes {
		fmt.Printf("%d:\t", size)
		shakerSort(random_arr[:size], "random")
	}

	fmt.Printf("\nSorted:\n")
	for _, size := range sliceSizes {
		fmt.Printf("%d:\t", size)
		shakerSort(sorted_arr[:size], "sorted")
	}

	fmt.Printf("\nInverted:\n")
	for _, size := range sliceSizes {
		fmt.Printf("%d:\t", size)
		shakerSort(inverted_arr[:size], "inverted")
	}
}

func generateRandomArray() []int {
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, 2000)
	for i := 0; i < 2000; i++ {
		arr[i] = rand.Intn(999999) + 1
	}
	return arr
}

func shakerSort(arr []int, _type string) {
	n := len(arr)
	left := 0
	right := n - 1
	swapped := true
	comparisons := 0
	swaps := 0

	for swapped {
		swapped = false

		for i := left; i < right; i++ {
			comparisons++
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swaps++
				swapped = true
			}
		}
		right--

		for i := right; i > left; i-- {
			comparisons++
			if arr[i] < arr[i-1] {
				arr[i], arr[i-1] = arr[i-1], arr[i]
				swaps++
				swapped = true
			}
		}
		left++
	}

	fmt.Printf("swaps: %-7d\t", swaps)
	fmt.Printf("comparisons: %-7d\t", comparisons)

	var theory_comparisons, theory_swaps int
	if _type == "random" {
		theory_comparisons = (n*n)/2
		theory_swaps = 3*(n*n-1)/4
	} else if _type == "sorted" {
		theory_comparisons = n*2 - 1
		theory_swaps = 0
	} else if _type == "inverted" {
		theory_comparisons = (n*n)/2
		theory_swaps = 3*(n*n-1)/2
	}

	fmt.Printf("[%-7d, %-7d]\n", theory_swaps, theory_comparisons)
}

