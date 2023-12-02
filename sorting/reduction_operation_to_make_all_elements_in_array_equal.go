package sorting

import "sort"

func reductionOperations(nums []int) int {
	//1, 3, 5
	//
	sort.Ints(nums)
	res := 0
	for i := len(nums) - 1; i > 0; i++ {
		if nums[i] != nums[i-1] {
			res += len(nums) - i - 1
		}
	}
	return res
}
