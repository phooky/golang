package main

import "fmt"

func isPalindrome(n int, base int) (bool) {
	digits := make([]int, 0, 22)
	for ; n > 0; n /= base {
		digits = append(digits,n % base)
	}
	l := len(digits)
	for i := 0; i < l/2; i++ {
		if digits[i] != digits[l-(i+1)] {
			return false
		}
	}
	return true
}

func isBoth(n int) (bool) {
	return isPalindrome(n,10) && isPalindrome(n,2)
}

func main() {
	sum := 0
	for i := 1; i < 1000000; i++ {
		if isBoth(i) {
			sum += i
			fmt.Println(i)
		}
	}
	fmt.Println(sum)
}