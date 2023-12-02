package array

import (
	"math"
	"sort"
)

func minPairSum(nums []int) int {
	//brute force - sort and 2 pointers
	sort.Ints(nums)
	res := math.MaxInt
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		res = max(res, nums[i]+nums[j])
	}
	return res
}
