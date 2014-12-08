
package main

import (
	"fmt"
	"math/big"
	"strings"
)

func main() {
	const precision = 2000
	biggest := 0
	biggestDem := 0
	for n := 2; n < 1000; n++ {
		r := big.NewRat(1, int64(n))
		rs := strings.TrimRight(r.FloatString(precision), "0")
		offset := 2
		for x := 1; (x < precision) && ((offset + x + 1 + x + 1) <= len(rs)); x++ {
			if rs[offset:offset+x+1] == rs[offset+x+1:offset+x+1+x+1] {
				size := len(rs[offset:offset+x+1])
				if size > biggest {
					biggest = size
					biggestDem = n
				}
				break
			}
		}
	}
	fmt.Println(biggest, biggestDem)
}
