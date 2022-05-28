package main
func longestPalindrome(s string) string {
    maxLength := 1
	startIndex := 0
	n := len(s)
	if (n < 2) {
		return s
	}
	for i := 1 ; i < n  ; i++ {
		leftIdx := i - 1
		rightIdx := i + 1 
		currentLen := 0
		currentIdx := 0
		for rightIdx < n && s[i] == s[rightIdx] {
			currentLen = rightIdx - i + 1 
			currentIdx = i
			rightIdx = rightIdx + 1
		}
		for leftIdx > -1 && s[leftIdx] == s[i] {
			currentLen = i - leftIdx + 1 
			currentIdx = leftIdx
			leftIdx = leftIdx - 1
		}
		for leftIdx > -1 && rightIdx < n && s[leftIdx] == s[rightIdx] {
			currentLen = rightIdx - leftIdx + 1
			currentIdx = leftIdx
			leftIdx = leftIdx - 1
			rightIdx = rightIdx + 1
		}
		if currentLen > maxLength {
			maxLength = currentLen 
			startIndex = currentIdx 
		}
	}
	return s[startIndex:startIndex+maxLength]

}

func GetLeftIdx (i int) int {
	if i > 0 {
		return i - 1
	}
	return -1
}

func GetRightIdx (i , len int) int {
	if i < len {
		return i + 1
	}
	return len
}
