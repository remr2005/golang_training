package main

import (
	"fmt"
	"math"
)

func DigPow(n, p int) int {
	res := 0
	for i := int(math.Log10(float64(n))); i >= 0; i-- {
		res += int(math.Pow(float64(n/int(math.Pow10(i))%10), float64(p)))
		p++
	}
	if res%n == 0 {
		return res / n
	}
	return -1
}

func main() {
	fmt.Println(DigPow(46288, 3))
}
