package main

import (
	"fmt"
	"math"
)

// FindMax to find the biggest value in array
func FindMax(arr []int, low, high int) int {
	if low == high {
		return arr[low]
	}

	mid := (low + high) / 2

	leftMax := FindMax(arr, low, mid)

	rightMax := FindMax(arr, mid+1, high)

	return int(math.Max(float64(leftMax), float64(rightMax)))
}

func main() {
	arr := []int{2, 5, 3, 8, 1, 6, 7, 4}
	maxValue := FindMax(arr, 0, len(arr)-1)
	fmt.Println("The maximum value is:", maxValue)
}
