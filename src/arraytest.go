package main

import "fmt"

func main() {
	//var base [10]int
	var base []int = make([]int, 10)
	base[4] = 5
	s1 := base[0:5]
	s2 := base[0:7]
	fmt.Println(" s1 len ",len(s1)," s1 cap ",cap(s1), " element 4 ",s1[4])
	fmt.Println(" s2 len ",len(s2)," s2 cap ",cap(s2), " element 4 ",s2[4])
	s2[4] = 8
	fmt.Println(" -- ")
	fmt.Println(" s1 len ",len(s1)," s1 cap ",cap(s1), " element 4 ",s1[4])
	fmt.Println(" s2 len ",len(s2)," s2 cap ",cap(s2), " element 4 ",s2[4])
	s2 = append(s2,4,4)
	fmt.Println(" -- ")
	fmt.Println(" s1 len ",len(s1)," s1 cap ",cap(s1), " element 4 ",s1[4])
	fmt.Println(" s2 len ",len(s2)," s2 cap ",cap(s2), " element 4 ",s2[4])
	s2[4] = 12
	fmt.Println(" -- ")
	fmt.Println(" s1 len ",len(s1)," s1 cap ",cap(s1), " element 4 ",s1[4])
	fmt.Println(" s2 len ",len(s2)," s2 cap ",cap(s2), " element 4 ",s2[4])
}