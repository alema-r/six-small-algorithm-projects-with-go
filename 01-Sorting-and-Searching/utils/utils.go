package utils

import (
	"fmt"
	"math/rand"
)

// MakeRandomArray creates an array of random numbers with numItems elements
// and all values less than max
func MakeRandomArray(numItems, max int) []int {
	randomArray := make([]int, numItems)
	for i := 0; i < numItems; i++ {
		randomArray[i] = rand.Intn(max)
	}
	return randomArray
}

// PrintArray prints either the complete array or numItems elements
func PrintArray(arr []int, numItems int) {
	if numItems > len(arr) {
		fmt.Println(arr)
	} else {
		fmt.Println(arr[:numItems])
	}
}

// CheckSorted prints to stdout whether the array is sorted or not.
func CheckSorted(arr []int) {
	if isSorted(arr) {
		fmt.Println("The array is sorted.")
	} else {
		fmt.Println("The array is not sorted.")
	}
}

// isSorted verify that an array is sorted.
func isSorted(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}
