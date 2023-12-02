package array

func restoreArray(adjacentPairs [][]int) []int {
	pairs := make(map[int][]int)
	for _, val := range adjacentPairs {
		pairs[val[0]] = append(pairs[val[0]], val[1])
		pairs[val[1]] = append(pairs[val[1]], val[0])
	}
	start := 0
	for key, val := range pairs {
		if len(val) == 1 {
			start = key
			break
		}
	}
	res := make([]int, len(adjacentPairs)+1)
	res[0] = start
	for i := 1; i < len(res); i++ {
		if len(pairs[res[i-1]]) == 1 || pairs[res[i-1]][1] == res[i-2] {
			res[i] = pairs[res[i-1]][0]
		} else {
			res[i] = pairs[res[i-1]][1]
		}
	}
	return res
}
