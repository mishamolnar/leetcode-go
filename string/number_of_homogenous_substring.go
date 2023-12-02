package string

func countHomogenous(s string) int {
	const (
		MOD = 1_000_000_007
	)
	char, count, res := '.', 0, 0
	for _, v := range s {
		if char == v {
			count++
		} else {
			count = 1
			char = v
		}
		res += count
		res %= MOD
	}
	return res
}
