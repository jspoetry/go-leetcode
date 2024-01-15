/*
 * @lc app=leetcode id=26 lang=golang
 *
 * [26] Remove Duplicates from Sorted Array
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	lastNumber := nums[0]
	lastIndex := 0

	for i, num := range nums {
		if num != lastNumber {
			lastNumber = num
			lastIndex++
			if lastIndex != i && i < len(nums) {
				nums[lastIndex] = num
			} 
		}
	}

	return lastIndex + 1
}
// @lc code=end

