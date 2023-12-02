package binary

import (
	"math"
	"strconv"
)

func MinimumOneBitOperations(n int) int {
	res := 0
	str := []rune(strconv.FormatInt(int64(n), 2))
	for i := range str {
		if str[i] == '1' {
			if i < len(str)-1 && str[i+1] == '1' {
				str[i] = '0'
				res += 1
			}
			break
		}
	}
	for i := range str {
		if str[i] == '1' {
			res += int(math.Pow(2.0, float64(len(str)-i))) - 1
		}
	}
	return res
}
