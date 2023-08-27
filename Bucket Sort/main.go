package main

import (
	"fmt"
	"sort"
)

func bucketSort(arr []float64, bucketSize int) []float64 {
	if len(arr) < 2 {
		return arr
	}

	// Find the max and min value
	minVal, maxVal := arr[0], arr[0]
	for _, num := range arr {
		if num < minVal {
			minVal = num
		}
		if num > maxVal {
			maxVal = num
		}
	}

	// Initialize
	bucketCount := int((maxVal-minVal)/float64(bucketSize) + 1)
	buckets := make([][]float64, bucketCount)

	// assign the data
	for i := 0; i < len(arr); i++ {
		idx := int((arr[i] - minVal) / float64(bucketSize))
		buckets[idx] = append(buckets[idx], arr[i])
	}

	// sort the buckets
	for i := 0; i < len(buckets); i++ {
		sort.Float64s(buckets[i])
	}

	// merge the bucket
	sortedIndex := 0
	for i := 0; i < len(buckets); i++ {
		for j := 0; j < len(buckets[i]); j++ {
			arr[sortedIndex] = buckets[i][j]
			sortedIndex++
		}
	}

	return arr
}

func main() {
	arr := []float64{0.12, 0.5, 0.46, 0.23, 0.89, 0.38, 0.72}
	fmt.Println("original array:", arr)
	sortedArr := bucketSort(arr, 5)
	fmt.Println("sorted array:", sortedArr)
}
