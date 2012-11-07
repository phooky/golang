package main

import "fmt"
import "math/big"

func nextFib(a,b *big.Int) (*big.Int) {
	next := new(big.Int)
	next.Add(a,b)
	return next
}

func tenToThe(n int) (*big.Int) {
	pow := big.NewInt(1)
	ten := big.NewInt(10)
	for ;n > 0; n-- {
		pow.Mul(pow,ten)
	}
	return pow
}

const digits int = 1000
		
func main() {
	minDigits := tenToThe(digits-1)
	a := big.NewInt(1)
	b := big.NewInt(1)
	n := 2
	for ;minDigits.Cmp(b) > 0; n++ {
		a, b = b, nextFib(a,b)
	}
	fmt.Println(n)
}
