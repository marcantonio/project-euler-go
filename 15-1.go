
package main

import (
	"fmt"
//	"os"
//	"strconv"
	"math"
)

func main() {
	var a [41][41]int
	for x := 0; x <= 40; x++ {
		for y := 0; y <= 40; y++ {
			switch {
			case y == 0:
				a[x][y] = 1
//				os.Stdout.Write([]byte(strconv.Itoa(a[x][y]) + " "))
			case y == x:
				a[x][y] = 1
//				os.Stdout.Write([]byte(strconv.Itoa(a[x][y]) + " "))
			case y  > x:
				continue
			default:
				a[x][y] = a[x-1][y-1] + a[x-1][y]
//				os.Stdout.Write([]byte(strconv.Itoa(a[x][y]) + " "))
			}
		}
//		os.Stdout.Sync()
//		fmt.Println()
	}
	
	// This mess is to get us down the triangle to the middle of the target row.  Every even row
	// (starting at 2) contants the answer for that row at it's center.  So, for a 1x1 grid the
	// count will hit at count == 1 which will be the 3rd row (x == 2).  The row is [1 2 1] and we
	// want the middle.  For a 2x2 grid count == 2, 5th row (x == 4), [1 4 6 4 1], the answer is 6.
	count := 0
	for x := 2; x < len(a); x++ {
		if (x % 2) == 0 {
			count++
		}
		// 20x20 grid.
		if count == 20 {
			// Half the size of x (which will always be odd here) + 1 is y.
			fmt.Println(a[x][int(math.Floor(float64(x/2)))])
			break
		}
	}
}
