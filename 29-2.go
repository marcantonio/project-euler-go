
package main

import (
	"fmt"
	"math/big"
)

func main() {
	arr := make([]*big.Int, 0)
	for a := 2; a <= 100; a++ {
		for b := 2; b <= 100; b++ {
			n := new(big.Int).Exp(big.NewInt(int64(a)), big.NewInt(int64(b)), nil)
			add := true
			for _, v := range arr {
				if n.Cmp(v) == 0 {
					add = false
				}
			}
			if add {
				arr = append(arr, n)
			}
		}
	}
	fmt.Println(len(arr))
}
