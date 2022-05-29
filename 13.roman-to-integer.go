package main
func romanToInt(s string) int {
	n := len(s)
	ans := 0
	for i := 0; i < n ; i++ {
		now := RomanToInt(s[i])
		after := 0 
		if i < n-1 {
			after = RomanToInt(s[i+1])
		}
		if now < after {
			ans += after - now
			i++
		} else {
			ans += now
		}
	}
	return ans
}

func RomanToInt(c byte) int {
	if c == 'I' {
		return 1
	}
	if c == 'V' {
		return 5
	}

	if c == 'X' {
		return 10
	}

	if c == 'L' {
		return 50
	}

	if c == 'C' {
		return 100
	}

	if c == 'D' {
		return 500
	}

	if c == 'M' {
		return 1000
	}
	return 0
}
