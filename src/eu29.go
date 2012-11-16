package main

//import "math"
import "fmt"

const MAX = 100

type Entry struct {
	base int
	power int
}

var entryMap map[Entry]bool
var marks [MAX+1]bool

func addEntries(n, max, power int) {
	for i := 2; i <= max; i++ {
		entry := Entry { n,i*power }
		entryMap[entry] = true
	}
}

func buildMaps(max int) {
	for i := 2; i <= max; i++ {
		for power, j := 1, i; j <= max; power, j = power+1, j*i {
			if marks[j] { continue }
			marks[j] = true
			addEntries(i,max,power)
		}
	}
}

func main() {
	entryMap = make(map[Entry]bool)
	marks = *new([MAX+1]bool)
	buildMaps(100)
	fmt.Println(len(entryMap))
}