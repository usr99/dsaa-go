package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := SliceToList([]int{2, 4, 3})
	l2 := SliceToList([]int{5, 6, 4})

	l3 := AddTwoNumbers(l1, l2)

	assert.Equal(t, ListToSlice(l3), []int{7, 0, 8})
}

// ListToSlice convert list to slice.
func ListToSlice(l *ListNode) (ret []int) {
	ret = make([]int, 0)
	for l != nil {
		ret = append(ret, l.Val)
		l = l.Next
	}

	return
}

// SliceToList convert slice to list.
func SliceToList(data []int) (l *ListNode) {
	settedHeader := false
	header, rear := &ListNode{}, &ListNode{}

	for _, val := range data {
		node := &ListNode{
			Val: val,
		}
		rear.Next = node

		if !settedHeader {
			header = rear.Next
			settedHeader = true
		}

		rear = rear.Next
	}

	return header
}
