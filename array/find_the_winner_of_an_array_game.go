package array

import "slices"

func getWinner(arr []int, k int) int {
	if k > len(arr) {
		return slices.Max(arr)
	}
	score := 0
	for {
		if arr[0] > arr[1] {
			score++
			arr = append(arr, arr[1])
			arr = append(arr[0:1], arr[2:]...)
		} else {
			score = 1
			arr = append(arr[1:], arr[0])
		}

		if score >= k {
			return arr[0]
		}
	}
}
