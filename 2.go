package main

import "fmt"

func main() {
	total := 0
	a, b := 1, 2
	for b < 4000000 {
		if (b % 2) == 0 {
			total += b
		}
		b, a = b + a, b
	}
	fmt.Println(total)
}
