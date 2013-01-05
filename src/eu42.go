package main

import "fmt"
import "os"
import "bufio"
import "strings"

var tris []int

func initTris() {
	maxtri := 26*30
	tris = make([]int,maxtri,maxtri)
	for i := 0; i < maxtri; i++ {
		tris[i] = (i*(i+1))/2
	}
}

func isTri(n int) (bool) {
	for _, val := range tris {
		if n == val {
			return true 
		}
	}
	return false
}

func val(s string) (int) {
	v := 0
	for _,c := range s {
		v += int(c - 'A')+1
	}
	return v
}


func main() {
	fmt.Println("SKY =",val("SKY"))
	initTris()
	file, err := os.Open("42-words.txt")
	if err != nil {
		panic("Could not open 42-words.txt")
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	count := 0
	for line,_ := reader.ReadString(','); line != ""; line,_ = reader.ReadString(',') {
		s := strings.Trim(line,"\",")
		if isTri(val(s)) {
			count++
		}
	}
	fmt.Println("count",count)
}