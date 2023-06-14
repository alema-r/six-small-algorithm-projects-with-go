package bubblesort

// BubbleSort sorts an array with the bubble sort algorithm
func BubbleSort(arr []int) {
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				swapped = true
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
	}
}
