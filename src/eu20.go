package main

import (
	"fmt"
	"math/big"
)

func fact(n int64) (*big.Int) {
	i := big.NewInt(1)
	for j := int64(2); j <= n; j++ {
		i.Mul(i,big.NewInt(j))
	}
	return i
}

func main() {
	var total int64 = 0
	for val := fact(100); val.Cmp(big.NewInt(0)) != 0; {
		_, mod := val.DivMod(val,big.NewInt(10),big.NewInt(0))
		total += mod.Int64()
	}
	fmt.Println(total)
}