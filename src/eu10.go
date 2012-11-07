package main

import (
	"fmt"
)


func markSieve(n int64, sieve []bool) {
	for m := n; m < int64(len(sieve)); m += n {
		sieve[m] = true
	}
}

func sumOfPrimes(max int64) (int64) {
	var sieve = make([]bool,max,max)
	var prime int64
	var sum int64
	for prime = 2 ; prime < max ; {
		sum += prime
		markSieve(prime,sieve)
		for prime++; prime < max && sieve[prime]; prime++ { }
	}
	return sum
}

func main() {
	fmt.Println("Prime sum under 10 ",sumOfPrimes(10))
	fmt.Println("Prime sum under 2m ",sumOfPrimes(2000000))
}