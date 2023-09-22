package main

import "strings"

func main() {
	lengthOfLongestSubstring("pwwkew")
}

func lengthOfLongestSubstring(s string) (l int) {
	var longest string
	var current string

	for _, char := range s {
		if index := strings.IndexRune(current, char); index != -1 {
			current = current[index+1:]
		}

		current += string(char)

		if len(current) > len(longest) {
			longest = current
		}
	}

	return len(longest)
}
