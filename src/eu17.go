package main

import "fmt"

func toText(n int) (string) {
	switch n {
	case 0: return "zero"
	case 1: return "one"
	case 2: return "two"
	case 3: return "three"
	case 4: return "four"
	case 5: return "five"
	case 6: return "six"
	case 7: return "seven"
	case 8: return "eight"
	case 9: return "nine"
	case 10: return "ten"
	case 11: return "eleven"
	case 12: return "twelve"
	case 13: return "thirteen"
	case 14: return "fourteen"
	case 15: return "fifteen"
	case 16: return "sixteen"
	case 17: return "seventeen"
	case 18: return "eighteen"
	case 19: return "nineteen"
	case 20: return "twenty"
	case 30: return "thirty"
	case 40: return "forty"
	case 50: return "fifty"
	case 60: return "sixty"
	case 70: return "seventy"
	case 80: return "eighty"
	case 90: return "ninety"
	}
	if n<100 {
		tens := (n/10) % 10
		ones := n % 10
		return toText(tens*10) + " " + toText(ones)
	}
	if n<1000 {
		hundreds := (n/100) % 10
		remainder := n % 100
		s := toText(hundreds) + " hundred"
		if remainder != 0 {
			s = s + " and "+toText(remainder)
		}
		return s
	}
	return "one thousand"
}

func countLetters(s string) (int) {
	count := 0
	for _,c := range s {
		if c != ' ' { count++ }
	}
	return count
}

func main() {
	count := 0
	for i := 1; i <= 1000; i++ {
		count += countLetters(toText(i))
	}
	fmt.Println(count)
	fmt.Println(toText(342),countLetters(toText(342)))
	fmt.Println(toText(115),countLetters(toText(115)))
	fmt.Println(toText(1000),countLetters(toText(1000)))
}