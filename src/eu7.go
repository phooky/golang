package main

import (
	"fmt"
)

const sieveSize = 10000000


func markSieve(n int64, sieve *[sieveSize]bool) {
	for m := n; m < sieveSize; m += n {
		sieve[m] = true
	}
}

func eu7(target int) (int64) {
	var sieve [sieveSize]bool
	var prime int64 = 2;
	for primeNum := 1; primeNum < target; primeNum++ {
		markSieve(prime,&sieve)
		for prime++; sieve[prime]; prime++ { }
	}
	return prime
}

func main() {
	fmt.Println("Prime 100 ",eu7(100))
	fmt.Println("Prime 6 ",eu7(6))
	fmt.Println("Prime 10001 ",eu7(10001))
}