package main
import "math"
var mem [][]int
func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	mem = make([][]int, n)
	for idx, value := range matrix {
		mem[idx] = make([]int, len(value))
		for index, _ := range value {
			dp(matrix, idx, index)
		}
	}
	if n == 0 {
		return 0
	}
	var ans int
	for idx, value := range mem[n-1] {
		if idx == 0 {
			ans = value
		} else if ans > value {
			ans = value
		}
	}
	return ans
}

func dp(matrix [][]int, i, j int) {

	if i == 0 {
		mem[i][j] = matrix[i][j]
	} else {
		mem[i][j] = Min(get(mem, i-1, j-1), get(mem, i-1, j), get(mem, i-1, j+1)) + matrix[i][j]
	}
}

func get(mem [][]int, i, j int) int {
	n := len(mem)
	m := len(mem[0])
	if i < 0 || i >= n || j < 0 || j >= m {
		return math.MaxInt
	}
	return mem[i][j]
}

func Min(a int, b int, c int) int {
	return min(min(a, b), c)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
