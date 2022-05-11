package main

import "testing"

func TestReverseList(t *testing.T) {
	var l1 *ListNode
	l1 = &ListNode{Val: 1}
	l1.Next = &ListNode{Val: 2}
	l1.Next.Next = &ListNode{Val: 3}
	l1.Next.Next.Next = &ListNode{Val: 4}
	l1.Next.Next.Next.Next = &ListNode{Val: 5}

	l2 := reverseList(l1)
	if l2.Val != 5 || l2.Next.Val != 4 || l2.Next.Next.Val != 3 || l2.Next.Next.Next.Val != 2 || l2.Next.Next.Next.Next.Val != 1 {
		t.Errorf("Expected 54321, got %d", l2.Val)
	} else {
		t.Logf("Expected 54321, got %d", l2.Val)
	}
}
