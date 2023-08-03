package main

import (
	"fmt"
	"reflect"
)

// Use for-range to check if the target exists in the array.
func ExistInArray(array []int, target int) bool {
	for _, value := range array {
		if target == value {
			return true
		}
	}
	return false
}

// Use for-loop to check if the target exists in the array.
func ExistInArray2(array []int, target int) bool {
	for i := 0; i < len(array); i++ {
		if target == array[i] {
			return true
		}
	}
	return false
}

// Check an element position in an array. If not, return -1
func CheckPosition(arr interface{}, target interface{}) int {
	array := reflect.ValueOf(arr)
	for i := 0; i < array.Len(); i++ {
		v := array.Index(i)
		if v.Interface() == target {
			return i
		}
	}
	return -1
}

// Method2 to check an element position in an array. If not, return -1
func checkArrayPosition(array []int, target int) int {
	for i := 0; i < len(array); i++ {
		if array[i] == target {
			return i
		}
	}
	return -1
}

// Delete duplicated elements in an unordered array.
func delDuplicate(array []int) []int {
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] == array[j] {
				array = append(array[:j], array[j+1:]...)
				j--
			}
		}
	}
	return array
}

// Method2 to delete duplicated elements in an unordered array.
func removeDuplicate(arr []int) []int {
	// We use map to store encountered values.
	encountered := map[int]bool{}
	result := []int{}
	for _, v := range arr {
		if !encountered[v] {
			encountered[v] = true
			result = append(result, v)
		}
	}
	return result
}

// Remove duplicated elements in an ordered array.
func Unique(array []int) []int {
	left := 0
	for right := 1; right < len(array); right++ {
		if array[left] != array[right] {
			left++
			array[left] = array[right]
		}
	}
	return array[:left+1]
}

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8}
	res := ExistInArray(array, 9)
	fmt.Println(res)
	res2 := ExistInArray2(array, 9)
	fmt.Println(res2)
	arrayPosition := CheckPosition(array, 5)
	fmt.Println(arrayPosition)
	arrayPosition2 := checkArrayPosition(array, 5)
	fmt.Println(arrayPosition2)
	array2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 1, 2}
	res3 := delDuplicate(array2)
	fmt.Println(res3)
	orderedArray := []int{1, 2, 2, 3, 3, 3, 4, 5, 5, 6}
	res4 := Unique(orderedArray)
	fmt.Println(res4)
}
