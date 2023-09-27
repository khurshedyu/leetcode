package main

func main() {
	isPalindrome(121)
	isPalindrome(-121)
	isPalindrome(10)
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	var reversed int
	var temp = x

	for temp != 0 {
		reversed = reversed*10 + temp%10
		temp /= 10
	}

	return reversed == x

}
