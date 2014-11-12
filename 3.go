
package main

import "fmt"

func main() {
	num := 600851475143
	accum := make([]int, 0)
	
	accum = append(accum, 1)
	factor(num, &accum)

	largest := 0
	for _, v := range accum {
		if v > largest {
			largest = v
		}
	}
	fmt.Println(accum)
	fmt.Println(largest)
}

func factor(num int, accum *[]int) []int {
	x := 0
	for x = 2; x <= num; x++ {
		if (num % x) == 0 {
			*accum = append(*accum, x)
			break
		}
	}
	if (num == 0) {
		return *accum
	}
	return factor(num/x, accum)
}
