package offer

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (head *ListNode) String() string {
	nextStr := "nil"
	if head.Next != nil {
		nextStr = head.Next.String()
	}
	return fmt.Sprintf("%d->%s", head.Val, nextStr)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
