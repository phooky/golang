package main

import "fmt"

var factTable [10]int

func initTable() {
	factTable[0] = 1
	for i := 1; i < 10; i++ {
		factTable[i] = factTable[i-1] * i
	}
}

// max soln will be 7 digits

func computeSums(n int) (int) {
	sum := 0
	for ; n > 0; n /= 10 {
		sum += factTable[n%10]
	}
	return sum
}

func main() {
	initTable()
	ans := 0
	for i := 10; i < 9999999; i++ {
		if computeSums(i) == i {
			fmt.Println(" - ",i)
			ans += i
		}
	}
	fmt.Println(ans)
}