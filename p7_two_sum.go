package main

import "math"

func TwoSumV1(pile []int, cible int) (int, int, bool) {
	dejavu := make(map[int]int)
	for i, v := range pile {
		deborde := (v > 0 && cible < math.MinInt+v) || (v < 0 && cible > math.MaxInt+v)
		if j, ok := dejavu[cible-v]; ok && !deborde {
			return pile[j], v, true
		}
		dejavu[v] = i
	}
	return 0, 0, false
}