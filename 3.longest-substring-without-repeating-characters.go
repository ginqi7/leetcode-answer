package main
import (
    "strings"
)
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}

	longest_str := ""
	longest_len := 0
	for i := 0 ; i < len(s); i++ {
		if strings.Contains(longest_str, string(s[i])) {
			index := strings.Index(longest_str, string(s[i]))
			longest_str = longest_str[index+1:len(longest_str)]
		} 
		longest_str = longest_str + string(s[i])
		longest_len = Max(longest_len, len(longest_str))
	}
	return longest_len
}


func Max(x, y int) int{
    if x < y {
        return y
    }
    return x
}
