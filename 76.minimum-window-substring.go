package main
import "math"
func minWindow(s string, t string) string {
	need := make(map[byte]int) 
	window := make(map[byte]int) 
	
	for idx, _ := range t {
		need = mapCount(need, t[idx])
	}
	
	
	left := 0
	right := 0
	valid := 0
	start := 0
	length := math.MaxInt
	
	for right < len(s) {
		c := s[right]
		right++
		
		_, ok := need[c]
		if ok {
			window = mapCount(window, c)
			if need[c] == window[c] {
				valid++
			}
		}
		
		for valid == len(need) {
			if right - left < length {
				length = right - left
				start = left
			}
			
			d := s[left]
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
	if length == math.MaxInt {
		return ""
	}

	return s[start:start+length]
}

func mapCount(aMap map[byte]int, c byte) map[byte]int{
	_, ok := aMap[c]
	if ok {
		aMap[c]++
	} else {
		aMap[c] = 1
	}
	return aMap
}	
