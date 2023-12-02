package array

import "strings"

func garbageCollection(garbage []string, travel []int) int {
	for i := 1; i < len(travel); i++ {
		travel[i] += travel[i-1]
	}
	return calculateGarbageType(garbage, travel, 'M') +
		calculateGarbageType(garbage, travel, 'P') +
		calculateGarbageType(garbage, travel, 'G')
}

func calculateGarbageType(garbage []string, travel []int, garbageType rune) int {
	garbageTime, travelTime := 0, 0
	for index, val := range garbage {
		garbageTime += strings.Count(val, string(garbageType))
		if index != 0 && strings.Contains(val, string(garbageType)) {
			travelTime = travel[index]
		}
	}
	return garbageTime + travelTime
}
