package reverse

func Reverse(input string) string {
	result := []rune(input)

	length := len(result)
	for i := 0; i < length/2; i++ {
		temp := result[i]
		result[i] = result[length-1-i]
		result[length-1-i] = temp
	}

	return string(result)
}
