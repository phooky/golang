package main

import "fmt"
import "libs/sieve"

const max = 1000000
var s = sieve.MakeSieve(max)

func testTrunc(p int) (bool) {
	if p < 10 {
		// fiat from problem def
		return false
	}
	// test right truncation
	for pright := p/10; pright > 0; pright /= 10 {
		if !s.IsPrime(int64(pright)) {
			return false
		}
	}
	// test left truncation
	for f := 10; f < p; f *= 10 {
		if !s.IsPrime(int64(p % f)) {
			return false
		}
	}
	return true
}

func main() {
	sum := 0
	for count, i := 0,10; count < 11; i++ {
		if s.IsPrime(int64(i)) && testTrunc(i) {
			count++
			sum += i
			fmt.Println(i)
		}
	}
	fmt.Println(sum)
}