package sorting

import "sort"

func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	//ans = math.min(num, ans + count(num)
	n := len(arr)
	count := make([]int, n+1)
	for _, val := range arr {
		count[min(val, n)]++
	}
	ans := 0
	for ind, numCount := range count {
		if numCount > 0 {
			ans = min(ind, ans+numCount)
		}
	}
	return ans
}

func maximumElementAfterDecrementingAndRearrangingSort(arr []int) int {
	sort.Ints(arr)
	curr := 1
	for _, val := range arr {
		if val > curr {
			curr++
		}
	}
	return curr
}
