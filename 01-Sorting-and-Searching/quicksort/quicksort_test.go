package quicksort

import (
	"testing"
)

func TestPartition(t *testing.T) {
	type testCase struct {
		initial, expected []int
		pivot_pos         int
	}

	tt := map[string]testCase{
		"array with three elements, all equal": {
			initial:   []int{0, 0, 0},
			expected:  []int{0, 0, 0},
			pivot_pos: 2,
		},
		"array with three elements, all distinct, not in order": {
			initial:   []int{0, 2, 1},
			expected:  []int{0, 1, 2},
			pivot_pos: 1,
		},
		"array sorted": {
			initial:   []int{1, 2, 3, 4, 5, 6},
			expected:  []int{1, 2, 3, 4, 5, 6},
			pivot_pos: 5,
		},
		"array not sorted": {
			initial:   []int{0, 5, 2, 3, 2, 6, 3},
			expected:  []int{0, 2, 3, 2, 3, 6, 5},
			pivot_pos: 4,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			arr := tc.initial
			p := partition(arr)
			if p != tc.pivot_pos {
				t.Errorf("expected %q as pivot position, got %q", tc.pivot_pos, p)
			}
			for i := range arr {
				if arr[i] != tc.expected[i] {
					t.Errorf("expected array to be %q, got %q", tc.expected, arr)
					break
				}
			}
		})
	}
}
