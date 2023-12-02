package sorting

import (
	"sort"
)

// m * n, because we maintain prevHeights always sorted
func largestSubmatrix(matrix [][]int) int {
	prevHeights := make([]BaseRow, 0)
	res := 0
	for i := range matrix {
		newHeights := make([]BaseRow, 0)
		seen := make([]bool, len(matrix[i]))
		for index := range prevHeights {
			if matrix[i][prevHeights[index].column] == 1 {
				newHeights = append(newHeights,
					BaseRow{prevHeights[index].height + 1, prevHeights[index].column})
			}
			seen[prevHeights[index].column] = true
		}
		for j := range matrix[i] {
			if !seen[j] && matrix[i][j] == 1 {
				newHeights = append(newHeights, BaseRow{1, j})
			}
		}
		for index := range newHeights {
			res = max(res, (index+1)*newHeights[index].height)
		}
		prevHeights = newHeights
	}
	return res
}

type BaseRow struct {
	height int
	column int
}

// m * n * log(n)
func largestSubmatrixWithSort(matrix [][]int) int {
	arr := make([]int, len(matrix[0]))
	res := 0
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				arr[j] = 0
			} else {
				arr[j]++
			}
		}
		calc := make([]int, len(matrix[0]))
		copy(calc, arr)
		res = max(res, biggestMatrix(calc))
	}
	return res
}

func biggestMatrix(arr []int) int {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})
	res := 0
	for index := range arr {
		res = max(res, (index+1)*arr[index])
	}
	return res
}

//[
//[1,1,0,0,1,0,1,1,0,1,1,1,1,0,1,1,1,0,1,1,1,0,1,0,0,1,1,1,1,1,0,1,0,1,1],
//[1,1,1,1,1,1,0,1,0,1,1,1,0,1,1,1,1,0,1,1,0,1,1,1,1,1,1,0,1,0,0,1,1,1,1],
//[1,1,1,0,1,1,1,1,1,0,0,0,1,1,1,0,1,1,1,1,1,1,1,1,1,0,1,1,1,0,1,1,1,1,0]
//]
