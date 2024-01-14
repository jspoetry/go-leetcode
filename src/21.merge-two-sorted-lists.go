/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	var prev, p1, p2 *ListNode

	p1 = list1
	p2 = list2
	head := p1

	for p2 != nil {
		switch {
			// no more values in first ll
			case p1 == nil:
				prev.Next = p2
				p2 = nil
			// value is lower
			case p2.Val <= p1.Val:
				if prev == nil {
					// update head
					prev = p2
					head = prev
				} else {
					prev.Next = p2
					prev = p2
				}
				p2 = p2.Next
				prev.Next = p1
			// move on
			case p2.Val > p1.Val:
				prev = p1
				p1 = p1.Next
		}
	}
		
	return head
}

// @lc code=end