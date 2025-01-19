package main

import "fmt"

func GetCount(str string) (count int) {
	for _, vow := range "aeiou" {
		for _, b := range str {
			if vow == b {
				count += 1
			}
		}
	}
	return count
}

func main() {
	fmt.Println()
}
