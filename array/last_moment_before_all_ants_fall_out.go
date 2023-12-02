package array

import "math"

func getLastMoment(n int, left []int, right []int) int {
	maxLeft, minRight := 0, math.MaxInt
	for _, v := range left {
		maxLeft = maxInt(maxLeft, v)
	}
	for _, v := range right {
		minRight = minInt(minRight, v)
	}
	return maxInt(maxLeft, n-minRight)
}

func maxInt(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func minInt(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
