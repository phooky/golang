package main

import "fmt"

var daysPerMonth [12]int = [...]int{ 
	31, // jan
	28, // feb
	31, // mar
	30, // apr
	31, // may
	30, // jun
	31, // jul
	31, // aug
	30, // sep
	31, // oct
	30, // nov
	31, // dec
}

func leapYear(year int) (bool) {
	return (year % 4 == 0) && ((year % 400 == 0) || (year % 100 != 0))
}

func daysInMonth(mon, year int) (int) {
	if (mon == 1 && leapYear(year)) { 
		return 29
	}
	return daysPerMonth[mon]
}

func main() {
	dayOfWeek := 1 // jan 1 1900 == monday
	total := 0
	// Take care of 1900. Project Euler, please stop
	// giving us misleading word problems.
	for mon := 0; mon < 12; mon++ {
		dayOfWeek = (dayOfWeek + daysInMonth(mon,1900)) % 7
	}
	for year := 1901; year < 2001; year++ {
		for mon := 0; mon < 12; mon++ {
			if (dayOfWeek == 0) { 
				total++
			}
			dayOfWeek = (dayOfWeek + daysInMonth(mon,year)) % 7
		}
	}
	fmt.Println(total)
}
	