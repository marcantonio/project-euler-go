package main

import "fmt"

func main() {
	primes := make([]int, 0)
	p := 0
	for i := 2; p != 10001; i++ {
		if isPrime(i) {
			primes = append(primes, i)
			p++
		}
	}
	fmt.Println(primes[len(primes)-1])
}

func isPrime(n int) bool {
	for x := 2; x < n; x++ {
		if (n % x) == 0 {
			return false
		}
	}
	return true
}
