package main

import (
	"fmt"
)

func main() {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("IV"))
	fmt.Println(romanToInt("IX"))
	fmt.Println(romanToInt("LVIII"))   // 58
	fmt.Println(romanToInt("MCMXCIV")) // 1994
}

func romanToInt(s string) int {
	var romanMap = map[string]int{
		"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000,
	}

	var num int
	for i := 0; i < len(s); i++ {

		if i < len(s)-1 {
			res := romanMap[string(s[i])] - romanMap[string(s[i+1])]
			if res < 0 {
				num += romanMap[string(s[i+1])] - romanMap[string(s[i])]
				i++
				continue
			}
		}

		num += romanMap[string(s[i])]
	}

	return num
}
