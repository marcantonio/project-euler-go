
package main

import (
	"fmt"
)

func main() {
	tri := 0
	for x := 1; ; x++ {
		tri += x
		
		factors := make([]int, 0)
		top := tri
		for n := 1; n < top; n++ {
			if (tri % n) == 0 {
				top = tri/n
				factors = append(factors, n, top)
			}
		}
		if len(factors) > 500 {
			for _, v := range factors {
				fmt.Print(v, " ")
			}
			fmt.Println()
			fmt.Println(len(factors))
			fmt.Println(tri)
			break
		}
	}
}
