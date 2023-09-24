package main

func twoSum(nums []int, target int) []int {
	hashmap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if v, ok := hashmap[target-nums[i]]; ok {
			return []int{v, i}
		}
		hashmap[nums[i]] = i
	}

	return []int{}
}
