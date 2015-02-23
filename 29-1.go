
package main

import (
	"fmt"
	"math/big"
)

func main() {
	arr := make([]*big.Int, 0)
	for a := 2; a <= 100; a++ {
		for b := 2; b <= 100; b++ {
			arr = append(arr, new(big.Int).Exp(big.NewInt(int64(a)), big.NewInt(int64(b)), nil))
		}
	}
	sort(&arr)

	count := 1
	for n := 0; n < len(arr)-1; n++ {
		if arr[n].Cmp(arr[n+1]) != 0 {
			count++
		}
	}

	fmt.Println(count)
}

func sort(arr *[]*big.Int) {
	for i := range *arr {
		for x := 0; x < len(*arr); x++ {
			// arr[i] < arr[x]
			if (*arr)[i].Cmp((*arr)[x]) == -1 {
				(*arr)[i], (*arr)[x] = (*arr)[x], (*arr)[i]
			}
		}
	}
}
