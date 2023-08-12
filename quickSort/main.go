package main

import (
	"fmt"
	"math/rand"
)

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[rand.Intn(len(arr))]

	low := make([]int, 0, len(arr))
	high := make([]int, 0, len(arr))
	pivotList := make([]int, 0, len(arr))

	for _, v := range arr {
		if v < pivot {
			low = append(low, v)
		} else if v > pivot {
			high = append(high, v)
		} else {
			pivotList = append(pivotList, v)
		}
	}

	low = quickSort(low)
	high = quickSort(high)

	low = append(low, pivotList...)
	low = append(low, high...)

	return low
}

func main() {
	arr := []int{64, 25, 12, 22, 11, 100, 200, 300, 13, 9, 1000, 5, 4, 2}
	fmt.Println("Original array:", arr)
	sortedArr := quickSort(arr)
	fmt.Println("Sorted array:", sortedArr)
}
