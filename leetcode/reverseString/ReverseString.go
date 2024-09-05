package reverseString

func ReverseString(s string) string {
	var newString string = ""
	array := []rune(s)

	for i := 0; i < len(s); i++ {

		newString = newString + string(array[len(s)-1-i])

	}

	return newString
}
func ReverseString1(s string) string {
	reversed := ""
	var length = len(s)
	for i := 0; i < length; i++ {
		reversed += string(s[length-1-i])

	}
	return reversed

}
