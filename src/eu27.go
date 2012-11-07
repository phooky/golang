package main

import "fmt"
import "libs/sieve"

var s *sieve.Sieve = sieve.MakeSieve(10000)

func consecutivePrimes(a, b int64) (int) {
	var i int64
	for i = 1; s.IsPrime(i*i + a*i + b); i++ { }
	return int(i)
}

func main() {
	maxA, maxB, maxCount := int64(0),int64(0),0

	for b := int64(0); b < 1000; b++ {
		if s.IsPrime(b) {
			for a := int64(-999); a < 1000; a++ {
				count := consecutivePrimes(a,b)
				if count > maxCount {
					maxA,maxB,maxCount = a,b,count
				}
			}
		}
				
	}
	fmt.Println(maxA,maxB,maxCount,maxA*maxB)
	//fmt.Println(1,41,consecutivePrimes(1,41))
	//fmt.Println(-79,1601,consecutivePrimes(-79,1601))
}
