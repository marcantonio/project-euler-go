
package main

import (
	"fmt"
)

type Pair struct {
	n int
	count int
}

func (p *Pair) String() string {
	return fmt.Sprintf("n: %v, count: %v", p.n, p.count)
}

func main() {
	p := new(Pair)
	
	for n := 1000000; n > 0; n-- {
		t := collatz(n)
		if t > p.count {
			p.count = t
			p.n = n
		}
	}
	fmt.Println(p)
}

func collatz(n int) int {
	even := func() { n = n / 2 }
	odd := func() { n = (3 * n) + 1 }
	
	count := 0
	for n != 1 {
		if (n % 2) == 0 {
			even()
		} else {
			odd()
		}
		count++
	}
	return count+1
}
