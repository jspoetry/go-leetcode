/*
 * @lc app=leetcode id=27 lang=golang
 *
 * [27] Remove Element
 */

// @lc code=start
func removeElement(nums []int, val int) int {
   uniqueNumIndex := -1

	 for i, num := range nums {
		if num != val {
			uniqueNumIndex++
			if i != uniqueNumIndex && uniqueNumIndex < len(nums) {
				nums[uniqueNumIndex] = num
			}
		}
	 }

	 return uniqueNumIndex + 1
}
// @lc code=end

