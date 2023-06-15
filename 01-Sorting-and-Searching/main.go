package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/alema-r/six-small-algorithm-projects-with-go/01-Sorting-and-Searching/bubblesort"
	"github.com/alema-r/six-small-algorithm-projects-with-go/01-Sorting-and-Searching/quicksort"
	"github.com/alema-r/six-small-algorithm-projects-with-go/01-Sorting-and-Searching/utils"
)

const (
	bubble = iota
	quick
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Get the number of items and maximum item value.
	var numItems, max, alg int
	fmt.Printf("# Items: ")
	fmt.Scanln(&numItems)
	fmt.Printf("Max: ")
	fmt.Scanln(&max)
	fmt.Printf("Algorithm (0: bubblesort, 1:quicksort): ")
	fmt.Scanln(&alg)

	// Make and display the unsorted array.
	arr := utils.MakeRandomArray(numItems, max)
	utils.PrintArray(arr, numItems)
	fmt.Println()

	// Sort with the algorithm chosen, do nothing if number is not 0 or 1
	switch alg {
	case bubble:
		bubblesort.BubbleSort(arr)
		fmt.Println("Using bubblesort.")
	case quick:
		quicksort.Quicksort(arr)
		fmt.Println("Using quicksort.")
	}
	// Display the result
	utils.PrintArray(arr, numItems)

	// Verify that it's sorted.
	utils.CheckSorted(arr)
}
