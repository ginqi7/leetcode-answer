package main
func corpFlightBookings(bookings [][]int, n int) []int {
    diff := make([]int, n)
	
	for _, val := range bookings {
		i := val[0] - 1
		j := val[1] - 1
		value := val[2]
		diff[i] += value
		if (j+1 < n) {
			diff[j+1] -= value
		}
	}
	
	return buildResult(diff)
}

func buildResult(diff []int) []int {
	n := len(diff)
	ans := make([]int, n)
	ans[0] = diff[0]
	for i := 1; i < n; i++ {
		ans[i] = ans[i-1] + diff[i]
	}
	return ans
}

