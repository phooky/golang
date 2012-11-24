package main

import "fmt"

func isEqual(n1 int, d1 int, n2 int, d2 int) (bool) {
	v1 := float32(n1)/float32(d1)
	v2 := float32(n2)/float32(d2)
	delta := v1 - v2
	if delta < 0 { delta = -delta }
	if delta < 0.00001 { fmt.Println(n1,d1,n2,d2) }
	return delta < 0.00000001 
}

func findCurious(n int, d int) (rn int, rd int) {
	rn = 1
	rd = 1
	for i := 1; i < 10; i++ {	
		if isEqual(n,d,n*10+i,d*10+i) || 
			isEqual(n,d,i*10+n,d*10+i) ||
			isEqual(n,d,n*10+i,i*10+d) ||
			isEqual(n,d,i*10+n,i*10+d) {
			rn = n
			rd = d
		}
	}
	return rn, rd
}

func main() {
	rn := 1
	rd := 1
	for d := 2; d < 10; d++ {
		for n := 1; n < d; n++ {
			nn, nd := findCurious(n,d)
			rn *= nn
			rd *= nd
		}
	}
	fmt.Println(rn,rd)
}
