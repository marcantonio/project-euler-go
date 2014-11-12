
package main

import (
	"fmt"
	"strconv"
)

func main() {
	largest := 0
	for a := 999; a > 10; a-- {
		for b := 999; b > 10; b-- {
			p := a * b
			if isPalindrome(strconv.Itoa(p)) {
				if p > largest {
					largest = p
				}
			}
		}
	}
	fmt.Println(largest)
}

func isPalindrome(str string) bool {
	for x, y := 0, len(str)-1; y > x; x, y = x+1, y-1 {
		if str[x] != str[y] {
			return false
		}
	}
	return true
}
