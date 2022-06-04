package main
func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	pre1 := 0
	pre2 := 1
	for i := 0 ; i < n; i ++ {
		temp := pre2
		pre2 = pre1 + pre2
		pre1 = temp
	}
	return pre1
}
