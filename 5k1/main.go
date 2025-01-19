package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func IpsBetween(start, end string) int {
	ip1 := strings.Split(start, ".")
	ip2 := strings.Split(end, ".")
	var res int64 = 0
	for i := 3; i >= 0; i-- {
		a, _ := strconv.ParseInt(ip1[i], 10, 64)
		b, _ := strconv.ParseInt(ip2[i], 10, 64)
		res += (b - a) * int64(math.Pow(256, float64(3-i)))
		fmt.Println(res, a, b)
	}
	return int(res)
}

func main() {
	fmt.Println(IpsBetween("10.0.0.0", "10.0.1.0"))
}
