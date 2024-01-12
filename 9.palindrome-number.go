/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	num := x
	reversed := 0

	// reverse number
	for num != 0 {
		// pick a last digit
		digit := num % 10
		// add digit to reversed
		reversed = reversed*10 + digit
		num /= 10
	}

	return reversed == x
}

// @lc code=end
