package string

func countPalindromicSubsequence(s string) int {
	//brute force: stop on each index and calculate number of charaters that occur in left and right
	//optimized: create an array of 24 and keep track of characters there
	right := [26]int{}
	for _, elem := range s {
		right[elem-'a']++
	}
	left := [26]int{}
	total := [26][26]bool{}
	for _, elem := range s {
		right[elem-'a']--
		for ind := range left {
			if left[ind] > 0 && right[ind] > 0 {
				total[elem-'a'][ind] = true
			}
		}
		left[elem-'a']++
	}
	res := 0
	for _, elem := range total {
		for _, b := range elem {
			if b {
				res++
			}
		}
	}
	return res
}
