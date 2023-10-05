package main

import (
	"testing"
)

func TestSliceRemove(t *testing.T) {
	s := []int{1,2,3,4,5}
	SliceRemove(s, 3)
	if exp := []int{1,2,3,5}; SliceEquality(exp, s) {
		t.Errorf("expected: %v, got: %v", exp, s)
	}
}

func TestReverseArrayPointer(t *testing.T) {
	arr := [3]int{1,2,3}
	ReverseArrayPointer(&arr)
	if exp := [3]int{3,2,1}; arr != exp {
		t.Errorf("expected: %v, got: %v", exp, arr)
	}
}