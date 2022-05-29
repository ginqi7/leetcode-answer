package main
func isValid(s string) bool {
	stack := []byte{}
	for i := 0; i < len(s) ; i++ {
		if len(stack) == 0 ||
			stack[len(stack) - 1] != findMatch(s[i]) {
			stack = append(stack, s[i])
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func findMatch(c byte) byte {
	if c == ')' {
		return '('
	}
	if c == '}' {
		return '{'
	}
	if c == ']' {
		return '['
	}
	return '0'
}
