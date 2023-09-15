package main

func maximumWealth(accounts [][]int) int {
	var maxWealth = 0
	for _, acc := range accounts {
		var sum = 0
		for _, wealth := range acc {
			sum += wealth
		}

		if sum > maxWealth {
			maxWealth = sum
		}
	}

	return maxWealth
}
