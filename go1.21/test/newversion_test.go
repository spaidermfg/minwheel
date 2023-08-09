package test

import (
	"fmt"
	"slices"
	"testing"
)

func TestMax(t *testing.T) {
	a := 7
	b := 8

	i := max(a, b)
	fmt.Println(i)
}

func TestMin(t *testing.T) {
	a := "go"
	b := "rust"

	s := min(a, b, "1python")
	fmt.Println(s)
}

func TestClear(t *testing.T) {
	a := []int32{1, 2, 3, 4, 5}
	clear(a)
	fmt.Println(a)
}

func TestSlices(t *testing.T) {
	a := []int32{1, 3, 2, 9, 4, 5, 7}
	index := slices.Index(a, 4)
	fmt.Println(index)

	indexFunc := slices.IndexFunc(a, func(i int32) bool {
		return i > 2
	})
	fmt.Println(indexFunc)

	names := []string{"mark", "ring", "queen", "king", "cot"}
	n, found := slices.BinarySearch(names, "mark")
	fmt.Println(n, found)
}
