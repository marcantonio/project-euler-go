
package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	bigun := big.NewInt(1)
	bigStr := fmt.Sprintf("%v", bigun.Exp(big.NewInt(2), big.NewInt(1000), nil))
	
	total := 0
	for _, v := range bigStr {
		i, _ := strconv.Atoi(string(v))
		total += i
	}
	fmt.Println(total)
}
