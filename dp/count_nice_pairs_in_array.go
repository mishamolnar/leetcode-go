package dp

const (
	MOD = 1_000_000_007
)

func countNicePairs(nums []int) int {
	//create map - difference -> count (not abs)
	//each time this diff occurs - add to result previous count
	counts := make(map[int]int)
	res := 0
	for _, num := range nums {
		diff := num - rev(num)
		res += counts[diff]
		res %= MOD
		counts[diff]++
	}
	return res
}

func rev(num int) int {
	res := 0
	for num > 0 {
		res *= 10
		res += num % 10
		num /= 10
	}
	return res
}
