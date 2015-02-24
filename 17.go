
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	numberMap := map[int]string{
		1:    "one",
		2:    "two",
		3:    "three",
		4:    "four",
		5:    "five",
		6:    "six",
		7:    "seven",
		8:    "eight",
		9:    "nine",
		10:   "ten",
		11:   "eleven",
		12:   "twelve",
		13:   "thirteen",
		14:   "fourteen",
		15:   "fifteen",
		16:   "sixteen",
		17:   "seventeen",
		18:   "eighteen",
		19:   "nineteen",
		20:   "twenty",
		30:   "thirty",
		40:   "forty",
		50:   "fifty",
		60:   "sixty",
		70:   "seventy",
		80:   "eighty",
		90:   "ninety",
		100:  "hundred",
		1000: "thousand",
	}
	
	count := 0
	for x := 1; x <= 1000; x++ {
		list := []string{}
		strNum := strconv.Itoa(x)
	done:
		for idx := range strNum {
			place := len(strNum) - idx
			
			digit, _ := strconv.Atoi(string(strNum[idx]))
			if digit == 0 {
				continue
			}
			switch place {
			case 1: // one's place
				list = append(list, numberMap[digit])
			case 2: // ten's place
				// special case for teens
				t, _ := strconv.Atoi(string(strNum[idx]) + string(strNum[idx+1]))
				if t < 20 {
					list = append(list, numberMap[t])
					break done
				}
				
				list = append(list, numberMap[digit*10])
			case 3: // hundred's place
				list = append(list, numberMap[digit] + " hundred")
				if (string(strNum[idx+1]) != "0") || (string(strNum[idx+2]) != "0") {
					list = append(list, "and")
				}
			case 4: // thousand's place
				list = append(list, numberMap[digit] + " thousand")
			}
		}
		
		fmt.Println(list)
		
		for i := range list {
			count += len(strings.Replace(list[i], " ", "", -1))
		}
	}
	fmt.Println(count)
}
