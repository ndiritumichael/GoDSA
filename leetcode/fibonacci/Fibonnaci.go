package fibonacci

func Fibonacci(n int) int {
	if n <= 0 {
		return 0
	}
	var fibMap *map[int]int = &map[int]int{}

	return getFib(n, *fibMap)
}

func getFib(n int, fibMap map[int]int) int {

	value, ok := fibMap[n]
	if ok {
		return value
	}
	var result int

	if n == 1 || n == 2 {
		result = 1
	} else {
		result = getFib(n-1, fibMap) + getFib(n-2, fibMap)
		fibMap[n] = result
	}
	return result
}
func GetFibWithoutRecursion(n int) int {
	var fibMap = make(map[int]int)

	if n <= 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	fibMap[1] = 1
	fibMap[2] = 1

	for i := 3; i <= n; i++ {
		fibMap[i] = fibMap[i-1] + fibMap[i-2]

	}

	return fibMap[n]

}
