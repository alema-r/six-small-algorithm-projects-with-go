package countingsort

import (
	"testing"

	"github.com/alema-r/six-small-algorithm-projects-with-go/01-Sorting-and-Searching/utils"
)

func TestCountingSort(t *testing.T) {
	type testCase struct {
		initial, expected []int
		max               int
	}

	tt := map[string]testCase{
		"array with one element": {
			initial:  []int{0},
			expected: []int{0},
			max:      0,
		},
		"array sorted": {
			initial:  []int{0, 1, 1, 2, 4},
			expected: []int{0, 1, 1, 2, 4},
			max:      4,
		},
		"array not sorted": {
			initial:  []int{4, 1, 0, 2, 5, 7},
			expected: []int{0, 1, 2, 4, 5, 7},
			max:      7,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			res := CountingSort(tc.initial, tc.max)
			for i := range res {
				if res[i] != tc.expected[i] {
					t.Errorf("the array is not sorted, got %q, expected %q", res, tc.expected)
					break
				}
			}
		})
	}
}

func TestCountingSortCustomer(t *testing.T) {
	type testCase struct {
		initial, expected []utils.Customer
		max               int
	}

	tt := map[string]testCase{
		"array with one element": {
			initial:  []utils.Customer{{Id: "C0", NumPurchases: 0}},
			expected: []utils.Customer{{Id: "C0", NumPurchases: 0}},
			max:      0,
		},
		"array sorted": {
			initial:  []utils.Customer{{Id: "C0", NumPurchases: 2}, {Id: "C1", NumPurchases: 3}, {Id: "C2", NumPurchases: 4}, {Id: "C3", NumPurchases: 5}},
			expected: []utils.Customer{{Id: "C0", NumPurchases: 2}, {Id: "C1", NumPurchases: 3}, {Id: "C2", NumPurchases: 4}, {Id: "C3", NumPurchases: 5}},
			max:      5,
		},
		"array not sorted": {
			initial:  []utils.Customer{{Id: "C0", NumPurchases: 5}, {Id: "C1", NumPurchases: 3}, {Id: "C2", NumPurchases: 3}, {Id: "C3", NumPurchases: 1}},
			expected: []utils.Customer{{Id: "C3", NumPurchases: 1}, {Id: "C1", NumPurchases: 3}, {Id: "C2", NumPurchases: 3}, {Id: "C0", NumPurchases: 5}},
			max:      5,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			res := CountingSortCustomer(tc.initial, tc.max)
			for i := range res {
				if res[i] != tc.expected[i] {
					t.Errorf("the array is not sorted, got %q, expected %q", res, tc.expected)
					break
				}
			}
		})
	}
}
