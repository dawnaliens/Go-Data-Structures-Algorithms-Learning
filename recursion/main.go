package main

import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	result := n * factorial(n-1)
	fmt.Printf("factorial(%d) = %d * factorial(%d) = %d\n", n, n, n-1, result)
	return result
}

func main() {
	factorial(10)
}
