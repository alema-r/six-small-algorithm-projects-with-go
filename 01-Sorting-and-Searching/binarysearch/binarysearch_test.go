package binarysearch

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	type testCase struct {
		arr              []int
		target, expected int
	}

	tt := map[string]testCase{
		"one element array, target present": {
			arr:      []int{0},
			target:   0,
			expected: 0,
		},
		"one element array, target not present": {
			arr:      []int{0},
			target:   1,
			expected: -1,
		},
		"five element array, target present": {
			arr:      []int{0, 1, 2, 3, 4},
			target:   4,
			expected: 4,
		},
		"five element array, target not present": {
			arr:      []int{0, 1, 3, 4, 5},
			target:   2,
			expected: -1,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			res, _ := BinarySearch(tc.arr, tc.expected)
			if res != tc.expected {
				t.Errorf("got %d, expected %d", res, tc.expected)
			}
		})
	}
}
