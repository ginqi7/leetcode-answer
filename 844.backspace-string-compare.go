package main
func backspaceCompare(s string, t string) bool {
    return getString(s) == getString(t)
}

func getString(s string) string {
	n := len(s)
	ans := make([]byte, n)
	if n == 0 {
		return  ""
	}
	count := 0
	for i := 0; i < n; i++ {
		if s[i] == '#' {
			if count > 0 {
				count--
			}
		} else {
			ans[count] = s[i]
			count++
		}
	}

	return string(ans[:count])
}
