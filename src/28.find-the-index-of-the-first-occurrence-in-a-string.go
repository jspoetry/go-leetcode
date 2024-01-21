/*
 * @lc app=leetcode id=28 lang=golang
 *
 * [28] Find the Index of the First Occurrence in a String
 */

// @lc code=start
func strStr(haystack string, needle string) int {
	p1 := 0
	p2 := -1 
	lastCharIdx := len(needle) - 1

	for p1 + lastCharIdx < len(haystack) {
		if p1 == p2 {
			return p1
		}

		if p2 != -1 {
			if haystack[p2] == needle[p2 - p1] {
				p2--
			} else {
				p2 = -1
				p1++
			}
		} else {
			if haystack[p1] == needle[0] {
				p2 = p1 + lastCharIdx
			} else {
				p1++
			} 
		}

	}

	return -1
}
// @lc code=end