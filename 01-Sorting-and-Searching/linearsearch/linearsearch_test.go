package linearsearch

import "testing"

func TestLinearSearch(t *testing.T) {
	type testCase struct {
		array                                   []int
		target, expectedIndex, expectedNumTests int
	}

	tt := map[string]testCase{
		"one-element array, target is present": {
			array:            []int{0},
			target:           0,
			expectedIndex:    0,
			expectedNumTests: 1,
		},
		"one-element array, target is not present": {
			array:            []int{0},
			target:           1,
			expectedIndex:    -1,
			expectedNumTests: 1,
		},
		"five-element array, target is present": {
			array:            []int{0, 3, 4, 2, 5},
			target:           2,
			expectedIndex:    3,
			expectedNumTests: 4,
		},
		"five-element array, target is not present": {
			array:            []int{0, 3, 4, 2, 5},
			target:           1,
			expectedIndex:    -1,
			expectedNumTests: 5,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			index, numTests := LinearSearch(tc.array, tc.target)
			if index != tc.expectedIndex {
				t.Errorf("Wrong index. Got %q, expected %q", index, tc.expectedIndex)
			}
			if numTests != tc.expectedNumTests {
				t.Errorf("Wrong number of tests. Got %q, expected %q", numTests, tc.expectedNumTests)
			}
		})
	}
}
