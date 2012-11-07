package main

import "fmt"

func nextInSeq(n int64) (int64) {
	if (n%2)==0 {
		return n/2
	}
	return 3*n + 1
}

const maxStart int64 = 1000000

var table [maxStart+1]int

func compute(start int64) (int) {
	if start <= 1 { return 1 }
	count := 1
	for n := start; n != 1; count++ {
		if n < start {
			count += table[n]
			break
		}
		n = nextInSeq(n)
	}
	table[start] = count
	return count
}

func main() {
	max := 0
	startOfMax := int64(0)
	for i := int64(1); i <= maxStart; i++ {
		if compute(i) > max {
			max = compute(i)
			startOfMax = i
		}
	}
	fmt.Println(startOfMax)
}

	