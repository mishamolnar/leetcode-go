package array

import "sort"

func maxCoins(piles []int) int {
	//9, 8, 7, 6, 5, 4, 3, 2, 1
	//9, 8, 1
	//7, 6, 2
	//5, 4, 3

	//approach: sort and add each second number
	sort.Ints(piles)
	res := 0
	for i := len(piles) - 2; i > len(piles)/3-1; i -= 2 {
		res += piles[i]
	}
	return res
}
