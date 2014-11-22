
package main

import (
	"fmt"
	"strconv"
	"math/big"
)

//func main() {
//	fmt.Println(factorial(10))
//	num := factorial(10)
//	numStr := strconv.Itoa(num)
//	
//	total := 0
//	for _, c := range(numStr) {
//		t, _ := strconv.Atoi(string(c))
//		total += t
//	}
//	fmt.Println(total)
//}
//
//func factorial(n int) int {
//	if n  == 1 {
//		return 1
//	}
//	return n * factorial(n - 1)
//}

func main() {
//	fmt.Println(mm.MulRange(1, 100))
	fmt.Println(factorial(big.NewInt(100)))
	num := factorial(big.NewInt(100))
	numStr := num.String()
	
	total := 0
	for _, c := range(numStr) {
		t, _ := strconv.Atoi(string(c))
		total += t
	}
	fmt.Println(total)
}

func factorial(n *big.Int) *big.Int {
	if n.Int64() == 1 {
		a := big.NewInt(1)
		return a
	}
	
	var tmp big.Int
	tmp.Set(n)
	return n.Mul(n, factorial(tmp.Sub(n, big.NewInt(1))))
}
