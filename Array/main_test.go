package main

import (
	"testing"
)

func TestExistInArray(t *testing.T) {
	array := []int{1, 2, 3, 4, 5}
	target := 3
	expected := true
	if result := ExistInArray(array, target); result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestExistInArray2(t *testing.T) {
	array := []int{1, 2, 3, 4, 5}
	target := 6
	expected := false
	if result := ExistInArray2(array, target); result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestCheckPosition(t *testing.T) {
	array := []int{1, 2, 3, 4, 5}
	target := 4
	expected := 3
	if result := CheckPosition(array, target); result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestCheckArrayPosition(t *testing.T) {
	array := []int{1, 2, 3, 4, 5}
	target := 7
	expected := -1
	if result := checkArrayPosition(array, target); result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestDelDuplicate(t *testing.T) {
	array := []int{1, 2, 2, 3, 4, 4, 5}
	expected := []int{1, 2, 3, 4, 5}
	if result := delDuplicate(array); !Equal(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestRemoveDuplicate(t *testing.T) {
	array := []int{1, 2, 2, 3, 4, 4, 5}
	expected := []int{1, 2, 3, 4, 5}
	if result := removeDuplicate(array); !Equal(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestUnique(t *testing.T) {
	array := []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}
	expected := []int{1, 2, 3, 4, 5}
	if result := Unique(array); !Equal(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
