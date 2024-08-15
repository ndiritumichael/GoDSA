package binarybridge

import "fmt"

func GetGap(N int) int {
	var first1Found bool = false
	var longestGap = 0
	var currentgap = 0

	binaryString := fmt.Sprintf("%b", N)

	for _, s := range binaryString {

		if s == '0' && first1Found == true {
			currentgap += 1
		} else if s == '1' {
			first1Found = true
			if currentgap > longestGap {
				longestGap = currentgap
			}
			currentgap = 0

		}

	}
	return longestGap

}
