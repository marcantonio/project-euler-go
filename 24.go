
package main

import (
	"fmt"
	"sort"
)

func main() {
	coll := "0123456789"
	accum := make([]string, 0)
	permute(coll, "", &accum)
	sort.Strings(accum)
	fmt.Println(accum[999999])
}

func permute(coll, pre string, accum *[]string) {
	if len(coll) == 2 {
		*accum = append(*accum, pre + coll, pre + string(coll[1]) + string(coll[0]))
		return
	}
	// Permute each rotation
	for x := 0; x < len(coll); x++ {
		permute(coll[1:], pre + string(coll[0]), accum)
		coll = coll[1:] + string(coll[0])
	}
}
