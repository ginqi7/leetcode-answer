package main
import (
    "bytes"
)
func convert(s string, numRows int) string {
	n := len(s)
	if numRows == 1 {
		return s
	}
	step := 2 * (numRows - 1)
	var ans bytes.Buffer
	for i :=0 ; i < numRows; i++ {
		for j := i; j < n; j = j+step {
			ans.WriteByte(s[j])
			if i != 0 && i != numRows -1  && (j + step - 2 * i) < n {
				midIdx :=  j + step - 2 * i    
				ans.WriteByte(s[midIdx])
			}
		}
	} 
    return ans.String()
}

/**
  * 0   4
  * 1 3 5
  * 2   6

  i * m + j * n + c
  1 * m + 0 * n + c = 3
  1 * m + 1 * n + c = 7
  1 * m + 2 * n + c = 11
  n = 4
  m = 

j + step - 2 * i)
2 * numRows - 2
  index + 2 * numRows - 2 - 2 *i
  index + 2 * (numRows -i -i)
 **/
