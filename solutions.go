package main

// solution to GOPL Exercises 4.3 - 4.7

// reverse an array
// edit in-place by passing a pointer to the array
func ReverseArrayPointer(arr *[3]int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	// return arr
}

// ex 4.4
func RotateSliceOnce(s []int, n int) {
	// keep a reserve of the n rotation-ed elements
	if n > len(s) {
		return
	}
	reserve := make([]int, n)
	copy(reserve, s[:n])
	copy(s, s[n:]) // copy the remainder into the slice
	copy(s[len(s)-n:], reserve)
}

// ex 4.5
// remove adjacent duplicates
// in-place function
func RemoveAdjacentDuplicates(s []string) []string {
	i := 0
	for _, v := range s {
		if s[i] == v {
			continue
		}
		i++
		s[i] = v
	}
	return s[:i+1]
}
