package distinctarray

import "testing"

type testvalues struct {
	arr  []int
	soln int
}

func TestDistinctArray(t *testing.T) {

	testCases := []struct {
		arr  []int
		soln int
	}{
		{[]int{2, 1, 1, 2, 3, 1}, 3},
		{[]int{1, 1, 1, 1, 1}, 1},
		{[]int{1, 2, 3, 4, 5}, 5},
		{[]int{}, 0},
		{[]int{-1, 0, 1, 2, -1, -4}, 5},
		{[]int{1, 1, 2, 2, 3, 3}, 3},
		{[]int{1, 2, 3, 1, 2, 3}, 3},
		{[]int{0, 0, 0, 0, 0}, 1},
		{[]int{100, 200, 100, 300}, 3},
		{[]int{-10, -20, -10, -30, -10}, 3},
	}

	for _, tc := range testCases {
		result := DistinctArraysolution(tc.arr)
		if result != tc.soln {
			t.Errorf("Expected %d, got %d for input %v", tc.soln, result, tc.arr)
		}
	}
}
