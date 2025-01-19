package main

import (
	"fmt"
	"math"
)

func TwoOldestAges(ages []int) [2]int {
	var res [2]int = [2]int{math.MinInt, math.MinInt}
	for _, val := range ages {
		if val >= res[0] {
			res[0] = val
		}
		if res[0] >= res[1] {
			res[0], res[1] = res[1], res[0]
		}
	}
	return res
}

func main() {
	fmt.Println(TwoOldestAges([]int{1, 23, 12323231, 12, 2, 3424}))
}
