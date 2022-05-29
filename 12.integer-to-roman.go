package main
import (
    "bytes"
)

func intToRoman(num int) string {
    var ans bytes.Buffer
	if num / 1000 > 0 {
		Write(&ans, 'M', num / 1000)
		num = num % 1000
	}
	if num / 900 > 0 {
		Write(&ans, 'C', 1)
		Write(&ans, 'M', 1)
		num = num % 900 
	}
	if num / 500 > 0 {
		Write(&ans, 'D', num / 500)
		num = num % 500
	}
	if num / 400 > 0 {
		Write(&ans, 'C', 1)
		Write(&ans, 'D', 1)
		num = num % 400
	}
	if num / 100 > 0 {
		Write(&ans, 'C', num / 100)
		num = num % 100
	}
	if num / 90 > 0 {
		Write(&ans, 'X', 1)
		Write(&ans, 'C', 1)
		num = num % 90
	}
	if num / 50 > 0 {
		Write(&ans, 'L', num / 50)
		num = num % 50
	}

	if num / 40 > 0 {
		Write(&ans, 'X', 1)
		Write(&ans, 'L', 1)
		num = num % 40
	}

	if num / 10 > 0 {
		Write(&ans, 'X', num / 10)
		num = num % 10
	}

	if num / 9 > 0 {
		Write(&ans, 'I', 1)
		Write(&ans, 'X', 1)
		num = num % 9
	}

	if num / 5 > 0 {
		Write(&ans, 'V', num / 5)
		num = num % 5
	}

	if num / 4 > 0 {
		Write(&ans, 'I', 1)
		Write(&ans, 'V', 1)
		num = num % 4
	}

	if num / 1 > 0 {
		Write(&ans, 'I', num / 1)
	}
	return ans.String()
}

func Write(ans *bytes.Buffer, c byte, num int) {
	for i := 0; i < num; i++ {
		ans.WriteByte(c)
	}
}
