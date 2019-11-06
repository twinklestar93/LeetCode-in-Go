package problem0539

import (
	"strconv"
	"sort"
	"math"
)

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func findMinDifference(timePoints []string) int {
	const N int = 60 * 24
	minutes := []int{}

	// convert timePoint into minutes
	for _, t := range(timePoints) {
		i1, _ := strconv.Atoi(t[:2])
		i2, _ := strconv.Atoi(t[3:])
		minutes = append(minutes, i1 * 60 + i2)
	}
	sort.Ints(minutes)
	var result int = math.MaxInt16
	for i:=0; i < len(timePoints); i++ {
		result = min(result, (N + minutes[(i+1) % len(timePoints)] - minutes[i]) % N)
	}
	return result
}


