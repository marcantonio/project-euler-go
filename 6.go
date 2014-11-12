package main

import "fmt"

func main() {
	sum := 0
	sumsum := 0
	for x := 1; x <=100; x++ {
		sum += x * x
		sumsum += x
	}
	fmt.Println((sumsum * sumsum) - sum)
}
