package main

import "fmt"

// BubbleSort
func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {

		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// exchange
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	data := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("Original Array:", data)
	BubbleSort(data)
	fmt.Println("Sorted Array:", data)
}
