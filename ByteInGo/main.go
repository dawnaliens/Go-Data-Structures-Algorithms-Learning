package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	tByte := []byte(t)
	sByte := []byte(s)
	sMap := map[byte]byte{}
	for i, b := range sByte {
		if _, ok := sMap[b]; !ok {
			sMap[b] = tByte[i]
		} else if sMap[b] != tByte[i] {
			return false
		}
	}
	return true
}

func main() {
	name := "Richard"
	bName := []byte(name)
	fmt.Println(bName)

	first := "apple"
	second := "egg"
	third := "add"
	res1 := isIsomorphic(first, second)
	res2 := isIsomorphic(third, second)
	res3 := isIsomorphic(third, first)
	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(res3)

}
