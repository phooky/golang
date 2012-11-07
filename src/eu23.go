package main

import "fmt"
import "libs/sieve"

const n = 28124
var s = sieve.MakeSieve(n)

type FactorTableEntry struct { factor, remainder int; abundant bool }
type FactorTable []FactorTableEntry

var tab = make(FactorTable,n)

func computePropDivisors(i int) (map[int] bool) {
	if i == 1 {
		return map[int] bool{ 1 : true }
	} else {
		fact := tab[i].factor
		rem := tab[i].remainder
		rv := computePropDivisors(rem)
		additions := make([]int,len(rv))
		j := 0
		for f,_ := range rv { 
			additions[j] = f*fact
			j++
		}
		for _,f := range additions {
			rv[f] = true
		}
		return rv
	}
	return nil
}

func computePropDivSum(i int) (int) {
	sum := 0
	divs := computePropDivisors(i)
	delete(divs,i)
	for f,_ := range divs {
		sum += f
	}
	return sum
}

func main() {
	tab[1] = FactorTableEntry{1,1,false}
	for i := 2; i < n; i++ {
		var j int
		for j = 2; j < i; j++ {
			if s.IsPrime(int64(j)) && (i % j == 0) {
				break
			}
		}
		tab[i] = FactorTableEntry{j,i/j,false}
		if tab[j].abundant || tab[i/j].abundant {
			tab[i].abundant = true
		} else {
			if computePropDivSum(i) > i { tab[i].abundant = true }
		}
	}
	var abundantSumTable [n]bool
	for i := 1; i < n; i++ {
		if tab[i].abundant {
			for j := i; j < n-i; j++ {
				if tab[j].abundant {
					abundantSumTable[i+j] = true
				}
			}
		}
	}
	sum := uint64(0)
	for i := 1; i < n; i++ {
		if !abundantSumTable[i] {
			sum += uint64(i)
		}
	}
        fmt.Println(sum)
}
