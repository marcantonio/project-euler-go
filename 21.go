
package main

import (
	"fmt"
)

func main() {
	total := 0
	for a := 10000; a > 0; a-- {
		b := d(a)
		if (b == 0) || (a == b) {
			continue
		}
		if a == d(b) {
			fmt.Println(a, b)
			total += a
		}
	}
	fmt.Println(total)
}

func d(a int) (t int) {
	t = 0
	for x := 1; x < a; x++ {
		if (a % x) == 0 {
			t += x
		}
	}
	return
}
