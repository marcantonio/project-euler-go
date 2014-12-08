
package main

import (
	"fmt"
	"math"
)

type Answer struct {
	a, b, count int
}

func (ans *Answer) coefficientProduct() int {
	return ans.a * ans.b
}

func main() {
	ans := new(Answer)
	for a := -999; a < 1000; a++ {
		for b := -999; b < 1000; b++ {
			count := check(a, b)
			if count > ans.count {
				ans.a = a
				ans.b = b
				ans.count = count
			}
		}
	}
	fmt.Println(ans.a, ans.b, ans.coefficientProduct(), ans.count)
}

func check(a, b int) int {
	var n int
	for n = 0; ; n++ {
		if !isPrime((n * n) + (a * n) + b) {
			break
		} 
	}
	return n
}

func isPrime(n int) bool {
	max := round(math.Sqrt(math.Abs(float64(n))))
	for x := 2; x <= max; x++ {
		if (n % x) == 0 {
			return false
		}
	}
	return true
}

func round(n float64) int {
	i, frac := math.Modf(n)
	if frac >= 0.5 {
		return int(i) + 1
	}		
	return int(i)
}
