package main

import "fmt"

func main() {
	for i := 1; i < 500; i++ {
		for j := 1; j < 500; j++ {
			c := 1000 - (i + j)
			if c*c == i*i + j*j {
				fmt.Println(i,j,c,i*j*c)
				return
			}
		}
	}
}
