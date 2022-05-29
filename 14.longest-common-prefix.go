func longestCommonPrefix(strs []string) string {
	idx := 0
	isLoop := true
	for isLoop {
		if idx >= len(strs[0]) {
			isLoop = false
			break
		}
		c := strs[0][idx]
		for i := 1; i < len(strs); i++ {
			if idx >= len(strs[i]) || c != strs[i][idx] {
				isLoop = false
				break
			}
			c = strs[i][idx]
		}
		if isLoop {
			idx ++
		}
	}
	return strs[0][0:idx]
}
