func generateParenthesis(n int) []string {
    ans := &[]string{}
	bfs(0, 0, n, "", ans)
	return *ans
}

func bfs (pos1 int, pos2 int, length int, item string, ans *[]string) {
	if pos1 == length && pos2 == length {
		(*ans) = append(*ans, item)
	}
	if pos1 < length {
		bfs(pos1+1, pos2, length, item + "(", ans)
	}

	if pos2 < pos1 {
		bfs(pos1, pos2+1, length, item + ")", ans)
	}
}
