package main
func checkInclusion(s1 string, s2 string) bool {
    need := make(map[byte]int)
    window := make(map[byte]int)
	
	for _, val := range s1 {
		need[byte(val)]++
	}
	

	
	left := 0
	right := 0
	valid := 0
	
	for right < len(s2) {
		c := s2[right]
		right++
		_, ok := need[c]
		if ok {
			window[c]++
			if window[c] == need[c]{
				valid++
			}
		}
		
		for right - left >= len(s1) {
			if valid == len(need) {
				return true
			}
			d := s2[left]
			left++
			_, ok := need[d]
			if ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}

	return false
}
