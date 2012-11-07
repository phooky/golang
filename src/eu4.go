package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(v int) (bool) {
	s := strconv.Itoa(v)
	l := len(s)
	for i := 0; i < (l+1)/2; i++ {
		if s[i] != s[(l-1)-i] {
			return false
		}
	}
	return true
}

func main() {
	i, j := 999, 999
	floor := 0
	for i > 99 {
		if i * j > floor {
			if isPalindrome(i*j) {
				floor = i * j
			}
		}
		if i == j {
			i--
			j = 999
		} else {
			j--
		}
	}
	fmt.Println(floor)
}