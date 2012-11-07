package main

import "fmt"
import "math/big"


func main() {
	num := big.NewInt(2)
	num.Exp(num, big.NewInt(1000), nil)
	s := fmt.Sprint(num)
	count := 0
	for _,c := range s {
		count += int(c) - '0'
	}
	fmt.Println(count)
}
