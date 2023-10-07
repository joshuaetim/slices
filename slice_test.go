package main

import (
	"testing"
)

func TestSliceRemove(t *testing.T) {
	s := []int{1,2,3,4,5}
	res := SliceRemove(s, 3)
	if exp := []int{1,2,3,5}; !SliceEquality(exp, res) {
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

func TestSliceIsNil(t *testing.T) {
	var s []int
	if !SliceIsNil(s) {
		t.Errorf("slice should be nil")
	}
}

func TestRotateSliceOnce(t *testing.T) {
	s := []int{0,1,2,3,4,5}
	n := 2

	exp := []int{2,3,4,5,0,1}
	if RotateSliceOnce(s, n); !SliceEquality(s, exp) {
		t.Errorf("expected: %v, got: %v", exp, s)
	}
}