package main

import "fmt"

func CopySlices() {
	a := []int{1, 2, 3}
	b := make([]int, 0, len(a))
	// to copy it is necessary that both slices are same length
	copy(b, a)
	fmt.Println(b)
}

func ReverseSlice(s []int) {
	// start from both ends and approach each other
	// subsitute opposite ends with each other
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	fmt.Println(s)
}

func SliceEquality(a, b []int) bool {
	// slice comparisions like this are shallow
	// fmt.Println(a == b) // fail
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// SliceIsNil checks if a slice is nil
func SliceIsNil(a []int) bool {
	// check 1: this is for empty, not nil
	// return len(a) == 0

	// check 2
	return a == nil
}

// CapacityVsArray shows how Go handles slices length and capacity
func CapacityVsArray() {
	a := make([]int, 0, 4)
	fmt.Println(a, len(a), cap(a))
	// due to 4 capacity, we can append
	a = append(a, 1, 2, 3)
	fmt.Println(a)

	// we can even append past capacity, more capacity will be created for us
	a = append(a, 4,5,6,7)
	fmt.Printf("a = %v, len=%d, cap=%d\n", a, len(a), cap(a))

	// capacity can NOT be less than length
	//a = make([]int, 19, 9)
	//fmt.Println(a, len(a), cap(a))
}

// SliceInPlaceChange shows how the same slice can be used in its manipulation
// without creating a new slice
func SliceInPlaceChange(s []string) []string {
	// example, removing empty strings from a slice
	out := s[:0]
	for i := range s {
		if s[i] != "" {
			out = append(out, s[i])
		}
	}
	return out
}

// SliceAsAStack implements a stack data structure using slices
func SliceAsAStack() {
	s := []int{}

	// push operation
	s = append(s, 1)
	s= append(s, 3)
	fmt.Println(s)

	// top operation
	s = append(s, 4)
	fmt.Println(s[len(s)-1])

	// pop operation
	fmt.Printf("before: %v\n", s)
	s = s[:len(s)-1]
	fmt.Printf("after: %v\n", s)

}

// SliceRemove removes an element from a slice using its index
func SliceRemove(s []int, i int) {
	// use the slice copy method

	//newS := s[:i] // don't do this; it returns a reference of the original slice
	// newS := make([]int, len(s[:i]))
	// copy(newS, s[:i])
	// newS = append(newS, s[i+1:]...)
	// fmt.Printf("result: %v\n", newS)

	// in-place method
	copy(s[i:], s[i+1:])
	// fmt.Printf("result (in-place): %v\n", s[:len(s)-1])
	// copy(s[:i], s[i+1:])
	// fmt.Printf("result: %v\n", s)
}

func main() {
	// fmt.Println(ReverseArrayPointer())
}
