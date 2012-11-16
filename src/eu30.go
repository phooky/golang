package main

import "fmt"

func fifthPowers() (powers [10]int) {
	for i := 0; i < 10; i++ {
		powers[i] = i * i * i * i * i
	}
	return powers
}

func testNumber(num int, fp *[10]int) (bool) {
	sum := 0
	for n := num;n > 0;n /= 10 {
		digit := n % 10
		sum += fp[digit]
	}
	return sum == num
}

func main() {
	powers := fifthPowers()
	sum := 0
	for i := 10; i < 9999999; i++ {
		if testNumber(i,&powers) {
			fmt.Println("Success",i)
			sum += i
		}
	}
	fmt.Println(sum)
}
