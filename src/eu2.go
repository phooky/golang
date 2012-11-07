package main

import "fmt"

const MAX int = 4000000

func nextFib(i1,i2 int) (int) {
	return i1 + i2
}

func main() {
	n1 := 1
	n2 := 2
	c := 0
	for n2 < MAX {
		if n2 % 2 == 0 {
			c += n2
		}
		n1,n2 = n2,nextFib(n1,n2)
	}
	fmt.Println(c)
}