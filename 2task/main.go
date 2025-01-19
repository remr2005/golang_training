package main

import (
	"fmt"
)

func Gimme(array [3]int) int {
	min, max := 0, 0
	for ind, val := range array {
		if val < array[min] {
			min = ind
		}
		if val > array[max] {
			max = ind
		}
	}
	return 3 - min - max
}

func main() {
	fmt.Println(Gimme([3]int{1, 2, 5}))
}
