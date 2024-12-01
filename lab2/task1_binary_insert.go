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
		binaryInsertionSort(random_arr[:size], "random")
	}

	fmt.Printf("\nSorted:\n")
	for _, size := range sliceSizes {
		fmt.Printf("%d:\t", size)
		binaryInsertionSort(sorted_arr[:size], "sorted")
	}

	fmt.Printf("\nInverted:\n")
	for _, size := range sliceSizes {
		fmt.Printf("%d:\t", size)
		binaryInsertionSort(inverted_arr[:size], "inverted")
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

func binaryInsertionSort(arr []int, _type string) {
	comparisons := 0
	swaps := 0

	for i := 1; i < len(arr); i++ {
		key := arr[i]
		insertPos := binarySearch(arr, 0, i-1, key)
		comparisons += i

		if insertPos != i {
			for j := i; j > insertPos; j-- {
				arr[j] = arr[j-1]
				swaps++
			}
			arr[insertPos] = key
		}
	}
	info(len(arr), swaps, comparisons, _type)
}

func binarySearch(arr []int, low, high, key int) int {
	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] == key {
			return mid
		} else if arr[mid] < key {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low
}

func info(n, swaps, comparisons int, _type string) {
	fmt.Printf("swaps: %-7d\t", swaps)
	fmt.Printf("comparisons: %-7d\t", comparisons)

	var theory_comparisons, theory_swaps int
	if _type == "random" {
		theory_comparisons = (n * (n - 1)) / 2
		theory_swaps = n*n
	} else if _type == "sorted" {
		theory_comparisons = (n * (n - 1)) / 2
		theory_swaps = 0
	} else if _type == "inverted" {
		theory_comparisons = (n * (n - 1)) / 2
		theory_swaps = (n * (n - 1)) / 2
	}

	fmt.Printf("[%-7d, %-7d]\n", theory_swaps, theory_comparisons)
}

