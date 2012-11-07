package main

import "fmt"
import "strings"
import "strconv"

var triangleSource string =
`75
95 64
17 47 82
18 35 87 10
20 04 82 47 65
19 01 23 75 03 34
88 02 77 73 07 63 67
99 65 04 28 06 16 70 92
41 41 26 56 83 40 80 70 33
41 48 72 33 47 32 37 16 94 29
53 71 44 65 25 43 91 52 97 51 14
70 11 33 28 77 73 17 78 39 68 17 57
91 71 52 38 17 14 91 43 58 50 27 29 48
63 66 04 68 89 53 67 30 73 16 69 87 40 31
04 62 98 27 23 09 70 98 73 93 38 53 60 04 23`

type triangle [][]int

func parseTriangle(source string) triangle {
	parsed := make([][]int,0,50)
	for _,lineStr := range strings.Split(source,"\n") {
		line := make([]int,0,50)
		for _,numStr := range strings.Split(lineStr," ") {
			num,_ := strconv.Atoi(strings.TrimSpace(numStr))
			line = append(line, num)
		}
		parsed = append(parsed,line)
	}
	return parsed
}

// Return the entry at the given row and column.
// Return 0 if out of bounds.
func (tri *triangle) at(row, col int) (int) {
	if (row < 0 || row >= len(*tri)) { return 0 }
	line := (*tri)[row]
	if (col < 0 || col >= len(line)) { return 0 }
	return line[col]
}

func (tri *triangle) leftParent(row,col int) (int) {
	return tri.at(row-1,col-1)
}

func (tri *triangle) rightParent(row,col int) (int) {
	return tri.at(row-1,col)
}

// Cascade 
func (tri *triangle) cascade() {
	for i := 0; i < len(*tri); i++ {
		for j := 0; j < len((*tri)[i]); j++ {
			var m int
			p1 := tri.leftParent(i,j)
			p2 := tri.rightParent(i,j)
			if (p1 > p2) { m = p1 } else { m = p2 }
			(*tri)[i][j] += m
		}
	}
}

func main() {
	tri := parseTriangle(triangleSource)
	tri.cascade()
	lastLine := tri[len(tri)-1]
	m := 0
	for _,val := range lastLine {
		if val > m { m = val }
	}
	fmt.Println(m)
}