package array

func getSumAbsoluteDifferences(nums []int) []int {
	res := make([]int, len(nums))
	prefix, suffix := 0, sum(nums)
	for index, val := range nums {
		res[index] = (index*val - prefix) + (suffix - (len(nums)-1-index)*val)
		prefix += val
		suffix -= val
	}
	return res
}

func sum(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		res += nums[i]
	}
	return res
}
