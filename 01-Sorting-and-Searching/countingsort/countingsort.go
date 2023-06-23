package countingsort

import (
	"github.com/alema-r/six-small-algorithm-projects-with-go/01-Sorting-and-Searching/utils"
)

// CountingSort implements the counting sort algorithm.
func CountingSort(arr []int, max int) []int {
	counts := make([]int, max+1)
	sorted := make([]int, len(arr))
	for _, v := range arr {
		counts[v]++
	}
	for i := 1; i < len(counts); i++ {
		counts[i] += counts[i-1]
	}
	for i := len(arr) - 1; i > -1; i-- {
		counts[arr[i]]--
		sorted[counts[arr[i]]] = arr[i]
	}
	return sorted
}

// CountingSortCustomer sorts a Customer array using counting sort.
func CountingSortCustomer(arr []utils.Customer, max int) []utils.Customer {
	counts := make([]int, max+1)
	sorted := make([]utils.Customer, len(arr))
	for _, v := range arr {
		counts[v.NumPurchases]++
	}
	for i := 1; i < len(counts); i++ {
		counts[i] += counts[i-1]
	}

	for i := len(arr) - 1; i > -1; i-- {
		counts[arr[i].NumPurchases]--
		sorted[counts[arr[i].NumPurchases]] = arr[i]
	}
	return sorted
}
