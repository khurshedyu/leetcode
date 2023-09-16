package main

func numberOfSteps(num int) int {
	var steps int = 0
	for {
		if num == 0 {
			return steps
		}

		var isOdd bool = num%2 != 0
		if isOdd {
			num = num - 1
			steps++
		} else {
			num = num / 2
			steps++
		}
	}
}
