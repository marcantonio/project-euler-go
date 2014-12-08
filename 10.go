package main

import (
	"fmt"
	"math"
)

func main() {
	sum := 0
	ch := make(chan bool, 4)
	for i := 2; i < 2000000; i++ {
		go isPrime(i, ch)
		
		if <- ch {
			sum += i
		}
	}
	fmt.Println(sum)
}

func isPrime(n int, ch chan bool) {
	max := round(math.Sqrt(math.Abs(float64(n))))
	for x := 2; x <= max; x++ {
		if (n % x) == 0 {
			ch <- false
			return
		}
	}
	ch <- true
}

func round(n float64) int {
	i, frac := math.Modf(n)
	if frac >= 0.5 {
		return int(i) + 1
	}		
	return int(i)
}
