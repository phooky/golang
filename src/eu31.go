package main

import "fmt"

var denominations []int = []int{ 1, 2, 5, 10, 20, 50, 100, 200 }

func countCombinations(total int, ds []int) (int) {
	if (total < 0) { return 0 }
	if (total == 0) { return 1 }
	combosum := 0
	for idx := len(ds)-1; idx >= 0; idx-- {
		d := ds[idx]
		for i := 1; total >= d*i; i++ {
			combosum += countCombinations(total - d*i,ds[:idx])
		}
	}
	return combosum
}

func main() {
	fmt.Println(countCombinations(200,denominations))
}
