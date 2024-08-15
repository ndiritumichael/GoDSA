package binarybridge

import "testing"

func TestGetGap(t *testing.T) {
	testCases := []struct {
		n   int
		gap int
	}{
		{8, 0},
		{15, 0},
		{22, 1},
		{5, 1},
		{6, 0},
		{1, 0},
		{0, 0},
		{31, 0},
		{1041, 5},
		{2147483647, 0},
	}

	for _, tc := range testCases {
		result := GetGap(tc.n)
		if result != tc.gap {
			t.Errorf("Expected gap %d for %d, got %d", tc.gap, tc.n, result)
		}
	}
}
