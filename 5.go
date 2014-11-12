package main

import "fmt"

func main() {
	top := 20
next:
	for x := 1; ; x++ {
		for y := 1; y != top; y++ {
			if (x % y) != 0 {
				continue next
			}
		}
		fmt.Println(x)
		break
	}
}
