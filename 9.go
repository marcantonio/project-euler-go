
package main

import "fmt"

func main() {
	for m := 2; m < 30; m++ {
		for n := 1; n < 30; n++ {
			triplet := calcPTriplet(m, n)
			val := triplet[0] + triplet[1] + triplet[2]
			if val == 1000 {
				fmt.Println("Found")
				fmt.Println(triplet)
				fmt.Println(triplet[0] * triplet[1] * triplet[2])
			}
		}
	}
}

func calcPTriplet(m, n int) []int {
	a := (m * m) - (n * n)
	b := 2 * m * n
	c := (m * m) + (n * n)
	
	return []int{a, b, c}
}
