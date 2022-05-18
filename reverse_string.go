package reverse_string

func ReverseString(input string) (output string) {
	r := []rune(input)
	ro := []rune{}
	for i := len(r) - 1; i >= 0; i-- {
		ro = append(ro, r[i])
	}
	output = string(ro)
	return output
}
