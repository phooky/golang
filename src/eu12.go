package main

import "fmt"
import "libs/sieve"

func triangleNumber(n int) (int) {
	return n * (n+1) / 2
}

var s *sieve.Sieve = sieve.MakeSieve(100000)

func countFactors(n int) (int) {
	count := 1
	for i := int(2); i <= n; i++ {
		if ! s.IsPrime(int64(i)) { continue }
		if n % i == 0 {
			x := 1
			for n % i == 0 {
				x++
				n /= i
			}
			count *= x
		}
	}
	return count
}

func main() {
	for i := 0; i < 1000000; i++ {
		nTri := triangleNumber(i)
		nCount := countFactors(nTri) 
		if nCount > 500 {
			fmt.Println(i,nTri,nCount)
			break
		}
	}
}