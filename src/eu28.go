package main

import "fmt"

func diagonals(prev, n int) (sum int, next int) {
	sum = 4*prev + 10*2*n
	next = prev + 4*2*n
	return sum, next
}

func diagsum(n int) (int) {
	sum := 1
	next := 1
	for i := 1; i <= (n-1)/2; i++ {
		var subsum int
		subsum, next = diagonals(next, i)
		fmt.Println(sum,subsum,next)
		sum += subsum
	}
	return sum
}

func main() {
	fmt.Println(diagsum(5))
	fmt.Println(diagsum(1001))
}
