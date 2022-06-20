package main
func climbStairs(n int) int {
	if n == 0 {
		return 0
	}
    if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	
	pre1 := 2
	pre2 := 1
	
	for i := 3; i <= n; i++ {
		current := pre1 + pre2
		pre2 = pre1
		pre1 = current
	}
	return pre1
}
