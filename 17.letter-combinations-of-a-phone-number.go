package main

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	digitsMap := []string{"", "abc", "def","ghi","jkl", "mno","pqrs","tuv","wxyz"}
	ans := []string{}
    bfs(0, len(digits), "", &ans, digitsMap, digits)
	return ans
}

func bfs(pos int, n int, str string, ans *[]string, digitsMap []string, digits string) {
	if pos == n {
		(*ans) = append(*ans, str) 
	} else {
		letters :=  digitsMap[digits[pos] - '0' - 1]
		for i := 0; i < len(letters) ; i++ {
			bfs(pos+1,n, str + string(letters[i]), ans, digitsMap, digits)
		}
	}
}
