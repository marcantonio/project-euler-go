
package main

import (
	"fmt"
)

func main() {
	dow := 1
	sundayCount := 0
	for year := 1900; year <= 2000; year++ {
		for month := 1; month <= 12; month++ {
			var days int
			switch month {
			case 4, 6, 9, 11:
				days = 30
			case 2:
				switch {
				case (year % 100) == 0:
					if (year % 400) == 0 {
						days = 29
					} else {
						days = 28
					}				
				case (year % 4) == 0:
					days = 29
				default:
					days = 28
				}
			default:
				days = 31
			}
			for day := 1; day <= days; day++ {
				if dow == 7 {
					if (year != 1900) && (day == 1) {
						sundayCount++
					}
					dow = 1
				} else {
					dow++
				}
			}
		}
	}
	fmt.Println(sundayCount)
}
