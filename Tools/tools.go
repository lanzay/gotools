package tools

import "unicode/utf8"

func Min(args ...int) int {
	
	min := args[0]
	for _, v := range args {
		if v < min {
			min = v
		}
	}
	return min
}

func Max(args ...int) int {
	
	max := args[0]
	for _, v := range args {
		if v > max {
			max = v
		}
	}
	return max
}

func Substring(s string, lenS int) string {
	
	by := []byte(s)
	posE := Min(lenS, len(s))
	for posE := posE; posE > 0; posE-- {
		if utf8.Valid(by[:posE]) {
			return string(by[:posE])
		}
	}
	return ""
	
	if len(by) > 1 {
		lenS = Min(lenS-1, len(s))
		if int(by[lenS-2]) >= 224 {
			lenS -= 2
		} else if int(by[lenS-1]) >= 192 && int(by[lenS-1]) < 224 {
			lenS -= 1
		}
		return s[0:lenS]
	}
	return s
}
