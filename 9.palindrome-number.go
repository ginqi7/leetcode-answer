package main
func isPalindrome(x int) bool {
    if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}

	if x % 10 == 0{
		return false
	}

	return x == reverse(x, 0)
}

func reverse(x int, ans int) int {
	if x == 0 {
		return ans
	}

	return reverse(x/10, ans * 10 + x%10)
}
