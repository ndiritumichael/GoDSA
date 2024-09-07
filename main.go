package main

import (
	"GoDSA/leetcode/binarybridge"
	"GoDSA/leetcode/fibonacci"
	"fmt"
	"time"
)

func main() {

	largestGap := binarybridge.GetGap(100)

	now := time.Now()
	fmt.Println("largest gap ", largestGap)

	fibNum := fibonacci.Fibonacci(50)
	fmt.Println("The fibonnaci number is ", fibNum)
	current := time.Since(now)
	print(" Done in about ", current.Milliseconds())
	fmt.Println("done")
}
