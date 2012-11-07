package main

import (
	"fmt"
)

func gcd(a, b int64) (int64) {
	if b == 0 { return a }
	return gcd(b, a%b)
}

func eu5(top int64) (int64) {
	var n, i int64
	n = 1
	for i = 1 ; i <= top; i++ {
		n = (n * i) / gcd(n,i)
	}
	return n
}

func main() {
	fmt.Println(eu5(20))
}