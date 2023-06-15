package quicksort

// partition splits an array in two partitions (Lomuto partition scheme)
// Lomuto partition scheme: scans the array with an index j and mantaining
// an index i such that, when the scan finishes, the elements from 0 to i-1
// (inclusive) are less than the pivot and the elements from i to j are
// greater than the pivot.
func partition(arr []int) int {
	// choose the last element as the pivot
	pivot := arr[len(arr)-1]
	i := -1
	for j := range arr {
		if arr[j] <= pivot {
			i++
			arr[j], arr[i] = arr[i], arr[j]
		}
	}
	return i
}

// Quicksort sorts an array with the quicksort algorithm
func Quicksort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	p := partition(arr)

	Quicksort(arr[:p])
	Quicksort(arr[p:])
}
