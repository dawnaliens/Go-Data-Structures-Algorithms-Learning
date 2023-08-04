package main

import (
	"fmt"
	"sort"
)

// Check the existence of a key in a map.
func ExistInMap(m map[int]string, key int) bool {
	_, ok := m[key]
	return ok
}

// Sort a map with slice.
func SortMap(m map[int]string) []string {
	if len(m) == 0 {
		return []string{}
	}
	var res []string
	for _, v := range m {
		res = append(res, v)
	}
	sort.Strings(res)
	return res
}

func main() {
	test := make(map[int]string)
	test[0] = "Richard"
	test[1] = "Tom"
	test[2] = "Jerry"
	test[3] = "Jack"
	test[10] = "Rose"
	test[20] = "Lily"
	test[12] = "Lucy"
	res := ExistInMap(test, 3)
	fmt.Println(res)

	res2 := SortMap(test)
	fmt.Println(res2)
}
