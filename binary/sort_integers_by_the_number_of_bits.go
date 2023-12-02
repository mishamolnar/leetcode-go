package binary

import (
	"math/bits"
	"sort"
)

func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		countI, countJ := bits.OnesCount(uint(arr[i])), bits.OnesCount(uint(arr[j]))
		if countJ == countI {
			return arr[i] < arr[j]
		} else {
			return countI < countJ
		}
	})
	return arr
}
