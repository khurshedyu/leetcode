package main

func runningSum(nums []int) []int {
	var sumOfNums = make([]int, len(nums)+1)

	for i := 1; i <= len(nums); i++ {
		sumOfNums[i] = sumOfNums[i-1] + nums[i-1]
	}

	return sumOfNums[1:]
}
