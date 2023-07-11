package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/alema-r/six-small-algorithm-projects-with-go/01-Sorting-and-Searching/binarysearch"
	"github.com/alema-r/six-small-algorithm-projects-with-go/01-Sorting-and-Searching/bubblesort"
	"github.com/alema-r/six-small-algorithm-projects-with-go/01-Sorting-and-Searching/countingsort"
	"github.com/alema-r/six-small-algorithm-projects-with-go/01-Sorting-and-Searching/linearsearch"
	"github.com/alema-r/six-small-algorithm-projects-with-go/01-Sorting-and-Searching/quicksort"
	"github.com/alema-r/six-small-algorithm-projects-with-go/01-Sorting-and-Searching/utils"
)

const (
	bubble = iota
	quick
	counting
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Get the number of items and maximum item value.
	var numItems, max, alg int
	var sorting bool
	fmt.Printf("# Items: ")
	fmt.Scanln(&numItems)
	fmt.Printf("Max: ")
	fmt.Scanln(&max)
	fmt.Printf("Do you want to search or sort? (search: 0, sort:1): ")
	fmt.Scanln(&sorting)

	if sorting {

		fmt.Printf("Algorithm (0: bubblesort, 1: quicksort, 2: countingsort): ")
		fmt.Scanln(&alg)

		// Sort with the algorithm chosen
		switch alg {

		case bubble:
			fmt.Println("Using bubblesort.")
			arr := utils.MakeRandomArray(numItems, max)
			utils.PrintArray(arr, numItems)
			bubblesort.BubbleSort(arr)
			utils.PrintArray(arr, numItems)
			utils.CheckSorted(arr)

		case quick:
			fmt.Println("Using quicksort.")
			arr := utils.MakeRandomArray(numItems, max)
			utils.PrintArray(arr, numItems)
			quicksort.Quicksort(arr)
			utils.PrintArray(arr, numItems)
			utils.CheckSorted(arr)

		case counting:
			fmt.Println("Using countingsort.")
			arr := utils.MakeRandomArrayCustomers(numItems, max)
			utils.PrintArrayCustomers(arr, numItems)
			arr = countingsort.CountingSortCustomer(arr, max)
			utils.PrintArrayCustomers(arr, numItems)
			utils.CheckSortedCustomer(arr)
		}

	} else {
		fmt.Printf("Algorithm (0: linear search, 1: binary search): ")
		fmt.Scanln(&alg)
		arr := utils.MakeRandomArray(numItems, max)
		if alg == 1 {
			quicksort.Quicksort(arr)
			utils.PrintArray(arr, numItems)
			utils.ArraySearchLoop(arr, binarysearch.BinarySearch)
		} else {
			utils.PrintArray(arr, numItems)
			utils.ArraySearchLoop(arr, linearsearch.LinearSearch)
		}

	}
}
