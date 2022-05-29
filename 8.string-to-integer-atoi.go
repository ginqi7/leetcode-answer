package main
import (
	"unicode"
	"math"
)
func myAtoi(s string) int {
	i := 0
	n := len(s)
	
	for i < n && isWhitespace(rune(s[i])) {
		i++ 
	}
	if i < n && isNegative(rune(s[i])) {
		return myAtoiHelper(s, i+1, n, 0, -1)
	}
	if i < n && isPositive(rune(s[i])) {
		return myAtoiHelper(s, i+1, n, 0, 1)
	}
	return myAtoiHelper(s, i, n, 0, 1)
}

func myAtoiHelper(s string, begin int, end int, ans int, c int) int {
	if begin >= end {
		return c * ans
	}
	if !isNum(rune(s[begin])) {
		return c * ans
	}
	
	ans = ans * 10 + int(s[begin]-'0')
	if (c == 1 && ans > math.MaxInt32) {
		return math.MaxInt32
	}
	if (c == -1 && -1 * ans < math.MinInt32) {
		return math.MinInt32
	}
	return myAtoiHelper(s, begin+1, end,  ans, c)
}

func isNum(c rune) bool{
	return unicode.IsNumber(c)
}

func isWhitespace(c rune) bool{
	return unicode.IsSpace(c)
}

func isNegative(c rune) bool {
	return c == '-'
}

func isPositive(c rune) bool {
	return c == '+'
}
