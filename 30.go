
package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	max := int(math.Pow(9, 5) * 4)
	var grandTotal int

	for x := 2; x < max; x++ {
		strNum := strconv.Itoa(x)
		var total int
		for _, v := range strNum {
			digit, _ := strconv.Atoi(string(v))
			total += int(math.Pow(float64(digit), 5))
		}
		if total == x {
			fmt.Println(total)
			grandTotal += total
		}
	}
	fmt.Println(grandTotal)
}
