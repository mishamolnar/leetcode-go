package dp

import (
	"sort"
)

func numFactoredBinaryTrees(arr []int) int {
	const MOD = 1_000_000_007
	sort.Ints(arr)
	indx := make(map[int]int)
	for i, val := range arr {
		indx[val] = i
	}
	var dp []int = make([]int, len(arr))
	for i := range dp {
		dp[i] = 1
	}
	for start, val := range arr {
		for i := start; i < len(arr); i++ {
			next := val * arr[i]
			index, ok := indx[next]
			if ok {
				add := int(int64(dp[i]*dp[start]) % MOD)
				if i != start {
					add *= 2
				}
				dp[index] += add
				dp[index] %= MOD
			}
		}
	}
	res := 0
	for _, val := range dp {
		res += val
		res %= MOD
	}
	return res
}
