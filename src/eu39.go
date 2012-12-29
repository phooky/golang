package main

import "fmt"

type Triple struct {
	a int
	b int
	c int
}

func gcd(a,b int) (int) {
	for ;b != 0; a,b = b,a%b {}
	return a
}

func perimeter(t Triple) (int) {
	return t.a + t.b + t.c
}

func generatePrimitiveTriple(m, n int) (Triple,bool) {
	var t Triple
	m2 := m*m
	n2 := n*n
	t.a = m2 - n2
	t.b = 2*m*n
	t.c = m2 + n2
	// presumes m > n
	if ((m-n)%2) == 0 {
		// m-n not odd
		return t,false
	}
	// check m, n coprime
	if (gcd(m,n) != 1) {
		return t,false
	}
	return t,true
}

func main() {
	var combomap [1001]int
	for m := 2; m < 100; m++ {
		for n := 1; n < m; n++ {
			t,b := generatePrimitiveTriple(m,n)
			p := perimeter(t)
			if p > 1000 {
				break
			}
			if b {
				for mp := p; mp <= 1000; mp += p {
					combomap[mp]++
				}
			}
		}
	}
	maxcombos := 0
	maxval := 0
	for i := 0; i <= 1000 ; i++ {
		if combomap[i] > maxcombos {
			maxval, maxcombos = i, combomap[i]
		}
	}
	fmt.Println(maxval, maxcombos)
}