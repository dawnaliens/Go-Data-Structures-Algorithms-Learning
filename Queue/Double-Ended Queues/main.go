package main

import (
	"errors"
	"fmt"
)

type DoubleQueue struct {
	slice []int
}

func (d *DoubleQueue) IsEmpty() bool {
	if len(d.slice) == 0 {
		return true
	}
	return false
}

func (d *DoubleQueue) AddLast(element int) {
	d.slice = append(d.slice, element)
}

func (d *DoubleQueue) AddFirst(element int) {
	if len(d.slice) == 0 {
		d.slice = append(d.slice, element)
	} else {
		d.slice = append([]int{element}, d.slice...)
	}
}

func (d *DoubleQueue) DeleteLast() error {
	if len(d.slice) == 0 {
		return errors.New("The Double Queue is empty.")
	} else {
		d.slice = d.slice[:len(d.slice)-1]
	}
	return nil
}

func (d *DoubleQueue) DeleteFirst() error {
	if len(d.slice) == 0 {
		return errors.New("The Double Queue is empty.")
	} else {
		d.slice = d.slice[1:]
	}
	return nil
}

func (d *DoubleQueue) DisplayFirst() (int, error) {
	if len(d.slice) == 0 {
		return 0, errors.New("The Double Queue is empty.")
	} else {
		return d.slice[0], nil
	}
}

func (d *DoubleQueue) DisplayLast() (int, error) {
	if len(d.slice) == 0 {
		return 0, errors.New("The Double Queue is empty.")
	} else {
		return d.slice[len(d.slice)-1], nil
	}
}

func main() {
	Double := &DoubleQueue{}
	Double.AddFirst(23)
	Double.AddLast(45)
	Double.AddFirst(10)
	fmt.Println(Double)
	fmt.Println(Double.IsEmpty())
	fmt.Println(Double.DisplayFirst())
	fmt.Println(Double.DisplayLast())
	Double.AddLast(2000)
	Double.AddFirst(100)
	fmt.Println(Double)
	Double.DeleteLast()
	fmt.Println(Double)
	Double.DeleteFirst()
	fmt.Println(Double)
}
