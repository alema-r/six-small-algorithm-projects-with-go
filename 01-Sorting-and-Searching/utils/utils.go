package utils

import (
	"fmt"
	"math/rand"
	"strconv"
)

const (
	sorted_str     = "The array is sorted."
	not_sorted_str = "The array is not sorted."
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
		fmt.Println(sorted_str)
	} else {
		fmt.Println(not_sorted_str)
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

type Customer struct {
	Id           string
	NumPurchases int
}

// MakeRandomArrayCustomers creates an array of random customers. It uses
// MakeRandomArray internally
func MakeRandomArrayCustomers(numItems, max int) []Customer {
	customerArray := make([]Customer, numItems)
	arr := MakeRandomArray(numItems, max)
	for i, v := range arr {
		customerArray[i] = Customer{Id: "C" + strconv.Itoa(i), NumPurchases: v}
	}
	fmt.Println(customerArray)
	return customerArray
}

// PrintArrayCustomers prints to stdout an array of Customer.
func PrintArrayCustomers(arr []Customer, numItems int) {
	if numItems > len(arr) {
		fmt.Println(arr)
	} else {
		fmt.Println(arr[:numItems])
	}
}

// CheckSortedCustomer prints to stdout whether an array of Customer is
// sorted or not.
func CheckSortedCustomer(arr []Customer) {
	if isSortedCustomer(arr) {
		fmt.Println(sorted_str)
	} else {
		fmt.Println(not_sorted_str)
	}
}

// isSortedCustomer checks if an array of Customer is sorted or not.
func isSortedCustomer(arr []Customer) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i].NumPurchases > arr[i+1].NumPurchases {
			return false
		}
	}
	return true
}
