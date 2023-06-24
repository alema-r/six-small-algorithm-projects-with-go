package linearsearch

// LinearSearch performs a linear search looking for target in arr.
// Returns the index of the target if it founds it (else -1) and
// the number of iterations that it took to find it.
func LinearSearch(arr []int, target int) (index, numTests int) {
	for i, v := range arr {
		if v == target {
			return i, i + 1
		}
	}
	return -1, len(arr)
}
