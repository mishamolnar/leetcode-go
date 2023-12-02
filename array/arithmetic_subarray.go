package array

import "slices"

func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	//brute force: copy an array from i to j, check if it is arithmetic
	//complexity: m * n
	res := make([]bool, 0, len(l))
	for i, left := range l {
		res = append(res, checkArithmetic(nums[left:r[i]+1]))
	}
	return res
}

func checkArithmetic(nums []int) bool {
	minValue, maxValue := slices.Min(nums), slices.Max(nums)
	if (maxValue-minValue)%(len(nums)-1) != 0 {
		return false
	}
	diff := (maxValue - minValue) / (len(nums) - 1)
	if diff == 0 {
		return true
	}
	set := make(map[int]bool, len(nums))
	for _, val := range nums {
		set[val] = true
	}
	for i := minValue; i <= maxValue; i += diff {
		if !set[i] {
			return false
		}
	}
	return true
}
