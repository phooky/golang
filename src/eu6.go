package main

import (
	"fmt"
	"math/big"
)

func sumOfSquares(a int64) (*big.Int) {
	n := new(big.Int)
	var i int64
	for i = 1; i <= a; i++ {
		n.Add(n,big.NewInt(i*i))
	}
	return n
}

func squareOfSum(a int64) (*big.Int) {
	n := new(big.Int)
	var i int64
	for i = 1; i <= a; i++ {
		n.Add(n,big.NewInt(i))
	}
	return n.Mul(n,n)
}

func main() {
	n := new(big.Int)
	n.Sub(squareOfSum(100),sumOfSquares(100))
	fmt.Println(n)
}