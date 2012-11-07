package main

import "fmt"
import "strings"
import "sort"
import "bufio"
import "os"

var nameList []string

func buildList(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		panic("problem opening "+path)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	rv := make([]string,0,6000)
	for line,_ := reader.ReadString(','); line != ""; line,_ = reader.ReadString(',') {
		rv = append(rv,strings.Trim(line,"\","))
	}
	return rv
}

func charValue(b rune) (int) {
	return int(b - 'A') + 1
}

func value(name string) (int) {
	total := 0
	for _,c := range name {
		total += charValue(c)
	}
	return total
}

func main() {
	nameList = buildList("names.txt")
	sort.Strings(nameList)
	total := uint64(0)
	for i := 0; i < len(nameList); i++ {
		total += uint64((i+1) * value(nameList[i]))
	}
	fmt.Println(total)
}