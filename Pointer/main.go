package main

import "fmt"

func main() {
	var a int = 23
	var b *int
	b = &a
	fmt.Println("The value of a is", a)
	fmt.Println("The value of *b is", *b)
	fmt.Println("The address of a is", &a)
	fmt.Println("The address of *b is", &(*b))
	*b = 45
	fmt.Println("After modifying *b, a =: ", a)
	fmt.Println("After modifying *b, *b = ", *b)
	fmt.Println("After modifying *b, &a = ", &a)
	fmt.Println("After modifying *b, &(*b) = ", &(*b))

	array := []int{23, 45, 99, 100, 200, 300, 400}
	var anotherArray *[]int
	anotherArray = &array
	fmt.Println("The value of array is", array)
	fmt.Println("The value of *anotherArray is", *anotherArray)
	fmt.Println("The address of array is", &array)
	fmt.Println("The address of *anotherArray is", &(*anotherArray))

	*anotherArray = []int{1, 2, 3, 4, 5}
	fmt.Println("After modifying *anotherArray, array =: ", array)
	fmt.Println("After modifying *anotherArray, *anotherArray = ", *anotherArray)
	fmt.Printf("After modifying *anotherArray, &array = %p\n", &array)
	fmt.Printf("After modifying *anotherArray, &(*anotherArray) = %p\n", &(*anotherArray))
}
