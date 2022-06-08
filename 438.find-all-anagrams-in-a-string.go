package main
func findAnagrams(s string, p string) []int {
	ans := []int{}
	need := make(map[byte]int)
	window := make(map[byte]int)
	for _, val := range p {
		c := byte(val)
		need[c]++
	}
	left := 0
	right := 0
	valid := 0
	
	for right < len(s) {
		c := s[right]
		right++
		_, ok := need[c]
		if ok {
			window[c]++
			if window[c] == need[c]{
				valid++
			}
		}
		
		for right - left >= len(p) {
			if valid == len(need) {
				ans = append(ans, left)
			}
			d := s[left]
			_,ok := need[d]
			if ok {
				if need[d] == window[d] {
					valid--
				}
				window[d]--
			}
			left++
		}
	}
	return ans


}
