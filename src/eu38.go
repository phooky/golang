package main

import "fmt"

func addToMap(digit int, dmap *[10]bool) (bool) {
	if dmap[digit] {
		return false
	}
	dmap[digit] = true
	return true
}

func try(n int) (int) {
	var dmap [10]bool
	total := make([]int,0,10)
	for f,digits := 1,9; digits > 0; f++ {
		factorStr := make([]int,0,10)
		for factor := f * n; factor > 0; factor /= 10 {
			fdigit := factor % 10
			if fdigit == 0 || !addToMap(fdigit, &dmap) {
				return 0
			}
			digits--
			factorStr = append(factorStr,fdigit)
		}
		for i := len(factorStr); i > 0; i-- {
			total = append(total,factorStr[i-1])
		}
	}
	var tint int
	for i := 0; i < len(total); i++ {
		tint = (tint * 10) + total[i]
	}
	return tint
}

func main() {
	var maxtry int
	for i := 1; i < 100000; i++ {
		v := try(i)
		if v > 0 {
			fmt.Println(i,v)
		}
		if v > maxtry {
			maxtry = v
		}
	}
	fmt.Println(maxtry)
}
