package main

import "fmt"

func findRepLen(numerator int, divisor int, digit int, table []int) (int) {
	remainder := numerator % divisor
	if remainder == 0 { return 0 }
	if table[remainder] != 0 {
		return digit - table[remainder]
	}
	table[remainder] = digit
	return findRepLen(remainder*10, divisor, digit+1, table)
}

func main() {
	maxCycle, maxValue := 0,0
	for i := 1; i < 1000; i++ {
		l := findRepLen(1,i,1,make([]int,i+1))
		if l > maxCycle {
			maxCycle, maxValue = l, i
		}
	}
	fmt.Println(maxValue,maxCycle)
}
