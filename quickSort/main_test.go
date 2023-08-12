package main

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestQuickSort(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	// Generate a slice of 100 random integers
	arr := make([]int, 100)
	for i := range arr {
		arr[i] = rand.Intn(1000) - rand.Intn(1000) // This will give both positive and negative numbers
	}

	sortedArr := quickSort(arr)

	// Verify if the slice is sorted
	if !sort.IntsAreSorted(sortedArr) {
		t.Errorf("quickSort did not sort the slice correctly")
	}

	// Let's ensure the sorted array has the same elements as the original one
	originalArrCopy := make([]int, len(arr))
	copy(originalArrCopy, arr)
	sort.Ints(originalArrCopy)
	if !reflect.DeepEqual(originalArrCopy, sortedArr) {
		t.Errorf("Sorted array elements don't match original array elements")
	}
}
