package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/alema-r/six-small-algorithm-projects-with-go/01-Sorting-and-Searching/bubblesort"
	"github.com/alema-r/six-small-algorithm-projects-with-go/01-Sorting-and-Searching/utils"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Get the number of items and maximum item value.
	var numItems, max int
	fmt.Printf("# Items: ")
	fmt.Scanln(&numItems)
	fmt.Printf("Max: ")
	fmt.Scanln(&max)

	// Make and display the unsorted array.
	arr := utils.MakeRandomArray(numItems, max)
	utils.PrintArray(arr, numItems)
	fmt.Println()

	// Sort and display the result.
	bubblesort.BubbleSort(arr)
	utils.PrintArray(arr, numItems)

	// Verify that it's sorted.
	utils.CheckSorted(arr)
}
