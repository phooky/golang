package main

import "fmt"
import "strconv"

type state struct {
	n int
	cur string
	idx int
}

func next(s *state) (int) {
	if len(s.cur) > s.idx {
		val := int(s.cur[s.idx]-'0')
		s.idx++
		return val
	}
	s.cur = strconv.Itoa(s.n)
	s.n++
	s.idx = 1
	return int(s.cur[0]-'0')
}
	
	
func main() {
	var s state
	s.cur = ""
	s.n = 1
	s.idx = 0
	total := 1
	for i := 1; i <= 1000000; i++ {
		n := next(&s)
		if i == 1 { total *= n }
		if i == 10 { total *= n }
		if i == 100 { total *= n }
		if i == 1000 { total *= n }
		if i == 10000 { total *= n }
		if i == 100000 { total *= n }
		if i == 1000000 { total *= n }
	}
	fmt.Println(total)
}
	
		
		
	