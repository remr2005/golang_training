package main

import (
	"strings"
)

func alphanumeric(str string) bool {
	if str == "" {
		return false
	}
	// fmt.Println(str)
	str = strings.ToLower(str)
	f := false
	for _, val := range str {
		for _, let := range "abcdefghijklmnopqrstuvwxyz1234567890" {
			if val == let {
				f = true
				break
			}
		}
		if f {
			f = false
			continue
		}
		return false
	}
	return true
}

func main() {
	// fmt.Println(alphanumeric(""))
}
