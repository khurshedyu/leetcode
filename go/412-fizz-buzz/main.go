package main

import "strconv"

func fizzBuzz(n int) []string {
	var answer []string
	for i := 1; i <= n; i++ {
		var divisibleBy3 = i%3 == 0
		var divisibleBy5 = i%5 == 0

		var str = ""

		if divisibleBy3 {
			str = "Fizz"
		}

		if divisibleBy5 {
			str += "Buzz"
		}

		if str == "" {
			str = strconv.Itoa(i)
		}

		answer = append(answer, str)
	}

	return answer
}
