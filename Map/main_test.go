package main

import (
	"reflect"
	"testing"
)

func TestExistInMap(t *testing.T) {
	testCases := []struct {
		m       map[int]string
		key     int
		expects bool
	}{
		{map[int]string{1: "one", 2: "two", 3: "three"}, 2, true},
		{map[int]string{1: "one", 2: "two", 3: "three"}, 4, false},
		{map[int]string{}, 1, false},
	}

	for _, tc := range testCases {
		result := ExistInMap(tc.m, tc.key)
		if result != tc.expects {
			t.Errorf("ExistInMap(%v, %v) = %v; expected %v", tc.m, tc.key, result, tc.expects)
		}
	}
}

func TestSortMap(t *testing.T) {
	testCases := []struct {
		m       map[int]string
		expects []string
	}{
		{map[int]string{1: "one", 2: "two", 3: "three"}, []string{"one", "three", "two"}},
		{map[int]string{}, []string{}},
	}

	for _, tc := range testCases {
		result := SortMap(tc.m)
		if !reflect.DeepEqual(result, tc.expects) {
			t.Errorf("SortMap(%v) = %v; expected %v", tc.m, result, tc.expects)
		}
	}
}
