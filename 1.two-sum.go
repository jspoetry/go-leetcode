/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	myMap := make(map[int]int)

	for index, num := range nums {
		secondNum := target - num
		secondIndex, hasTarget := myMap[secondNum]

		if hasTarget {
			return []int{secondIndex, index}
		} else {
			myMap[num] = index
		}
	}

	return []int{}
}

// @lc code=end
