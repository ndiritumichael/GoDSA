package reverseString

func ReverseString(s string) string {
	var newString string = ""
	array := []rune(s)

	for i := 0; i < len(s); i++ {

		newString = newString + string(array[len(s)-1-i])

	}

	return newString
}
