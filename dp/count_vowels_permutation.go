package dp

func CountVowelPermutation(n int) int {
	var dp = make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 5)
	}
	for i := range dp[0] {
		dp[0][i] = 1
	}
	res, MOD := 0, 1_000_000_007
	next := [][]int8{{1}, {0, 2}, {0, 1, 3, 4}, {2, 4}, {0}}
	for i := range dp {
		for j := range dp[i] {
			if i < len(dp)-1 {
				for _, k := range next[j] {
					dp[i+1][k] += dp[i][j]
					dp[i+1][k] %= MOD
				}
			} else {
				res += dp[i][j]
				res %= MOD
			}
		}

	}
	return res
}
