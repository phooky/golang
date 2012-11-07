package main

import "fmt"

// The number of permutations of a given number of elements
func permutationCount(n int) (uint64) {
	if n <= 1 { return 1 }
	return uint64(n)*permutationCount(n-1)
}

func getNthElement(digits []int, n int) ([]int, int) {
	v := digits[n]
	copy(digits[n:],digits[n+1:])
	digits = digits[:len(digits)-1]
	return digits, v
}

func main() {
	digits := make([]int,10,10)
	for i := 0; i < 10; i++ {
		digits[i] = i
	}
	var permutation [10]int
	index := uint64(1000000 - 1)
	for i:=0; i<10; i++ {
		subperm := permutationCount(9-i)
		nth := index / subperm
		index -= nth*subperm
		digits, permutation[i] = getNthElement(digits,int(nth))
	}
	for i:=0; i<10; i++ {
		fmt.Print(permutation[i])
	}
	fmt.Println()
}