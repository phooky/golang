package main

import "fmt"

func valueOf(digits []int) (int) {
	sum := 0
	for _,digit := range digits {
		sum *= 10
		sum += digit
	}
	return sum
}

func selectIdx(digits []int, idx int) (value int, remaining []int) {
	remaining = make([]int, len(digits)-1)
	copy(remaining[:idx],digits[:idx])
	copy(remaining[idx:],digits[idx+1:])
	return digits[idx],remaining
}

func search3(digits []int, target int, v2 int, v1 int) (bool) {
	for i := 0; i < len(digits); i++ {
		value, remaining := selectIdx(digits, i)
		newV1 := (v1*10) + value
		if len(remaining) == 0 && (newV1*v2) == target {
			return true
		}
		if search3(remaining,target,v2,newV1) {
			return true
		}
	}
	return false
}

func search2(digits []int, target int, v2 int) (bool) {
	for i := 0; i < len(digits); i++ {
		value, remaining := selectIdx(digits, i)
		newV2 := (v2*10) + value
		if search3(remaining,target,newV2,0) {
			return true
		}
		if search2(remaining,target,newV2) {
			return true
		}
	}
	return false
}

func search(digits []int, target int) (sum int) {
	sum = 0
	for i := 0; i < len(digits); i++ {
		value, remaining := selectIdx(digits, i)
		newTarget := (target*10) + value
		if search2(remaining,newTarget,0) {
			sum += newTarget
			fmt.Println(newTarget)
		}
		sum += search(remaining,newTarget)
	}
	return sum
}

func main() {
	digits := [...]int{ 1, 2, 3, 4, 5, 6, 7, 8, 9 }
	fmt.Println("Sum:",search(digits[:],0))
}