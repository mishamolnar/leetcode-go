package sorting

import (
	"math"
	"slices"
	"sort"
)

func eliminateMaximum(dist []int, speed []int) int {
	for i, distance := range dist {
		dist[i] = int(math.Ceil(float64(distance) / float64(speed[i])))
	}
	slices.SortFunc()
	sort.Ints(dist)
	for i, val := range dist {
		if val <= i {
			return i
		}
	}
	return len(dist)
}
