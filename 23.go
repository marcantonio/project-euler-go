
package main

import (
	"fmt"
	"math"
)

func main() {
	abundantNums := make([]int, 0)
	for x := 12; x <= 28123; x++ {
		if isAbundant(x) {
			abundantNums = append(abundantNums, x)
		}
	}
	
	abundantSums := make(map[int]bool)
	for n := 0; n < len(abundantNums); n++ {
		for m := n; m < len(abundantNums); m++ {
			total := abundantNums[n] + abundantNums[m]
			if _, ok := abundantSums[total]; !ok {
				abundantSums[total] = true
			}
		}
	}
	total := 0
	for x := 1; x <= 28123; x++ {
		if _, ok := abundantSums[x]; !ok {
			total += x
		}
	}
	fmt.Println(total)
}

func isAbundant(num int) bool {
	divisors := 0
	
	for n := 1; n <= num/2; n++ {
		if (num % n) == 0 {
			divisors += n
		}
	}
	
	if divisors > num {
		return true
	}
	return false
}

func round(n float64) int {
	i, frac := math.Modf(n)
	if frac >= 0.5 {
		return int(i) + 1
	}		
	return int(i)
}
