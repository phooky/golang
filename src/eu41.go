package main

import "fmt"
import "math"

type node struct {
	sofar int
	remaining []int
}

func IsPrime(p int) (bool) {
	top := int(math.Ceil(math.Sqrt(float64(p))))
	if p%2 == 0 { return false }
	for i := 3; i <= top; i +=2 {
		if p%i == 0 { return false }
	}
	return true
}

func findMax(n node, m int) (int) {
	if len(n.remaining) == 0 {
		// don't bother to eval if under current max
		if n.sofar > m && IsPrime(n.sofar) {
			return n.sofar
		}
		return m
	}
	// explore remainder of tree
	for i := 0; i < len(n.remaining); i++ {
		n2 := n
		n2.sofar = (n2.sofar*10) + n.remaining[i]
		n2.remaining = make([]int,0,len(n.remaining)-1)
		for j := 0; j < len(n.remaining); j++ {
			if j != i {
				n2.remaining = append(n2.remaining,n.remaining[j])
			}
		}
		m = findMax(n2,m)
	}
	return m
}

func main() {
	var n node
	n.sofar = 0
	n.remaining = []int{1,2,3,4,5,6,7,8,9}
	m := 0
	for ;m == 0; {
		m = findMax(n,m)
		n.remaining = n.remaining[:len(n.remaining)-1]
	}
	fmt.Println(m)
}
