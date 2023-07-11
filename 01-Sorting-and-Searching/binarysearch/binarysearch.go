package binarysearch

// BinarySearch performs a logarithmic search on array arr, looking for target.
// Returns the index if target is found else -1 and the number of tests made to
// find it.
func BinarySearch(arr []int, target int) (index, numTests int) {
	l := 0
	r := len(arr) - 1
	numTests = 0
	for l <= r {
		numTests++
		m := (l + r) / 2
		if arr[m] < target {
			l = m + 1
		} else if arr[m] > target {
			r = m - 1
		} else {
			return m, numTests
		}
	}
	return -1, numTests
}
