package main

import "fmt"

func nextFactor(v, min int64) (fact int64) {
	for fact := min; fact < v; fact++ {
		if v % fact == 0 {
			return fact
		}
	}
	return v
}

func main() {
	var fact, v int64
	fact = 2
	v = 600851475143
	for fact < v {
		fact = nextFactor(v,fact)
		v /= fact
		fmt.Println("Factor: ",fact)
	}
	fmt.Println("Maximum prime factor: ",fact)
}