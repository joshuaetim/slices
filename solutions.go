package main

// solution to GOPL Exercises 4.3 - 4.7

func ReverseArrayPointer(arr *[3]int) {
	for i, j := 0, len(arr) - 1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	// return arr
}