
package main

import (
	"fmt"
	"math/big"
)

func main() {
	fib(big.NewInt(1), big.NewInt(1), 1)
}

func fib(a, b *big.Int, c int) {
	c++
	if len(b.String()) == 1000 {
		fmt.Println(c)
		return
	}
	fib(b, a.Add(a, b), c)
}
