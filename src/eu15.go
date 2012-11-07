package main

import "fmt"
import "math/big"

// Let's see:
// The problem is equivalent to how many ways we
// can stuff X balls into Y buckets, where a "ball"
// is a downstroke in the grid.
// This is a multiset coefficient with k = balls
// and n = buckets.

func countCombos(balls, buckets int64) (*big.Int) {
	// work out n and k for binomial coefficient
	n := balls + buckets - 1
	k := balls

	cumul := big.NewInt(1)
	for i := n; i > (n-k); i-- {
		cumul.Mul(cumul,big.NewInt(i))
	}
	for i := k; i >= 1; i-- {
		cumul.Div(cumul,big.NewInt(i))
	}
	return cumul
}

func main() {
	fmt.Println(countCombos(2,3))
	fmt.Println(countCombos(10,11))
	fmt.Println(countCombos(20,21))
}