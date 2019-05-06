package leetcode

// ListNode defined the singly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// AddTwoNumbers Description:
// You are given two non-empty linked lists representing two non-negative
// integers. The digits are stored in reverse order and each of their nodes
// contain a single digit. Add the two numbers and return it as a linked list.
// You may assume the two numbers do not contain any leading zero, except the number 0 itself.
//
// Example:
//
//	Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
//	Output: 7 -> 0 -> 8
//	Explanation: 342 + 465 = 807.

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	settedHeader := false
	header, rear := &ListNode{}, &ListNode{}

	action := func(sum int) {
		if rear.Next != nil {
			sum += rear.Next.Val
		}
		current, front := SplitNumber(sum)

		node := &ListNode{
			Val: current,
		}
		rear.Next = node

		if !settedHeader {
			header = rear.Next
			settedHeader = true
		}

		rear = rear.Next

		if front != 0 {
			rear.Next = &ListNode{
				Val: front,
			}
		}

	}

	for l1 != nil && l2 != nil {
		action(l1.Val + l2.Val)
		l1 = l1.Next
		l2 = l2.Next
	}
	for l1 != nil {
		action(l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		action(l2.Val)
		l2 = l2.Next
	}

	return header
}

// SplitNumber split number to current and front.
func SplitNumber(val int) (int, int) {
	return val % 10, val / 10
}
