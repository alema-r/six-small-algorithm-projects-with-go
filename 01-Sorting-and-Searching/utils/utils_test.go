package utils

import "testing"

func TestMakeRandomArray(t *testing.T) {
	type testCase struct {
		numItems, max int
	}

	tt := map[string]testCase{
		"numItems: 5, max: 1": {
			numItems: 5,
			max:      1,
		},
		"numItems: 10, max: 10": {
			numItems: 10,
			max:      10,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			arr := MakeRandomArray(tc.numItems, tc.max)
			if len(arr) != tc.numItems {
				t.Errorf("Invalid length, expected %q, got %q", tc.numItems, len(arr))
			}
			for _, v := range arr {
				if v >= tc.max || v < 0 {
					t.Errorf("Got %q, expected values only in the interval [0,%q)", v, tc.max)
				}
			}
		})
	}
}

func ExampleCheckSorted() {
	CheckSorted([]int{0, 1, 2, 2, 3})
	CheckSorted([]int{0, 2, 1, 3, 4})
	// Output:
	// The array is sorted.
	// The array is not sorted.
}

func TestIsSorted(t *testing.T) {
	type testCase struct {
		arr    []int
		sorted bool
	}

	tt := map[string]testCase{
		"empty array, sorted": {
			arr:    []int{},
			sorted: true,
		},
		"one element array, sorted": {
			arr:    []int{1},
			sorted: true,
		},
		"general sorted array, sorted": {
			arr:    []int{0, 4, 5, 5, 9},
			sorted: true,
		},
		"general unsorted array, not sorted": {
			arr:    []int{1, 2, 10, 9, 11},
			sorted: false,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			res := isSorted(tc.arr)
			if res != tc.sorted {
				t.Errorf("Got %t, expected %t", res, tc.sorted)
			}
		})
	}
}
