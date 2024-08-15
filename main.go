package main

import (
	"GoDSA/binarybridge"
	"fmt"
)

func main() {

	largestGap := binarybridge.GetGap(100)

	fmt.Println("largest gap ", largestGap)

}
