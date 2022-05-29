package main

// 给定一个链表，判断链表中是否有环，用O(1)内存解决。
// 如有有环，返回入环的那个节点

// 解决思路：
// 使用141题的思路，即快慢双指针，看指针是否相遇判断有无环
// 有环，则使用第三个指针从头节点走，慢指针从相遇节点走
// 两指针第一次相遇节点极为入口节点
func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	fast := head
	slow := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			temp := head
			for temp != slow {
				temp = temp.Next
				slow = slow.Next
			}
			return temp
		}
	}
	return nil
}
