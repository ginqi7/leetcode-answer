package main
func findOcurrences(text string, first string, second string) []string {
    splits := strings.Split(text, " ")
	n := len(splits)
	ans := []string{}
	for i := 0; i <= n - 3; i++ {
		if splits[i] == first && splits[i+1] == second {
			ans = append(ans, splits[i+2])
		}
	}
	return ans;
}
