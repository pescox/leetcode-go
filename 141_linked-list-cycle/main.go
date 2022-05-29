package main

// 给定一个链表，判断链表中是否有环，用O(1)内存解决。

// 解题思路：快慢双指针，快指针每次走2步，慢指针每次走1步，
// 如果存在环
// 环的大小小于等于n，快慢指针一定会相遇
func main() {

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	fast := head
	slow := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
