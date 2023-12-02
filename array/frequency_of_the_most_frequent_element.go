package array

import (
	"fmt"
	"sort"
)

func maxFrequency(nums []int, k int) int {
	//1, 2, 4, 5
	//1, 3, 7, 12
	//1, 4, 12, 20
	sort.Ints(nums)
	prefix := make([]int, len(nums)+1)
	pre := 0
	for index, val := range nums {
		pre += val
		prefix[index+1] = pre
	}
	left, res := 0, 0
	fmt.Println(nums)
	fmt.Println(prefix)
	for index, val := range prefix {
		for (index-left+1)*nums[index]-(val-prefix[left]) > k {
			left++
		}
		fmt.Printf("left %d , right %d \n", left, index)
		res = max(res, index-left+1)
	}
	return res
}
