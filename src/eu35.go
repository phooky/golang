package main

import "fmt"
import "libs/sieve"

const max = 1000000
var s = sieve.MakeSieve(max)

func countDigits(n int) (int) {
	digits := 0
	for ; n > 0; n /= 10 {
		digits += 1
	}
	return digits
}

func rotate(n int, digits int) (int) {
	p := n % 10
	for ; digits > 1; digits-- {
		p *= 10
	}
	return p + (n/10)
}

func checkCircularPrime(n int) (bool) {
	digits := countDigits(n)
	for d := digits; d > 0; d-- {
		n = rotate(n,digits)
		if !s.IsPrime(int64(n)) { return false }
	}
	return true
}
	
func main() {
	count := 0
	for i := 2; i < max; i++ {
		if checkCircularPrime(i) {
			count++
			fmt.Println(i)
		}
	}
	fmt.Println(count)
}