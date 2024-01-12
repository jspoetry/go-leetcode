/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefixEnd := 0
	unmatch := false
	firstStr := strs[0]

	// loop over strings before
	for prefixEnd != len(firstStr) && !unmatch {
		nextChar := firstStr[prefixEnd]

		for _, str := range strs {
			if prefixEnd >= len(str) || str[prefixEnd] != nextChar {
				unmatch = true
				break
			}
		}

		if !unmatch {
			prefixEnd++
		}
	}

	return string(firstStr[:prefixEnd])
}

// @lc code=end