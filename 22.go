
package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"sort"
)

func main() {
	charMap := make(map[rune]int)
	for r, i := 'A', 1; r <= 'Z'; r, i = r+1, i+1 {
		charMap[r] = i
	}
	
	data, err := ioutil.ReadFile("./p022_names.txt")
	if err != nil {
		panic(err)
	}
	
	names := strings.Split(string(data), ",")
	for i := range names {
		names[i] = strings.Replace(names[i], "\"", "", -1)	
	}
	sort.Strings(names)
	
	total := 0
	for i, name := range names {
		nameTotal := 0
		for _, r := range name {
			nameTotal += charMap[r]
		}
		total += nameTotal * (i + 1) 
	}
	fmt.Println(total)
}
