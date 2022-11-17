package offer


// 剑指 Offer 4. 从尾到头打印链表
// 输入一个链表的头节点，按链表从尾到头的顺序返回每个节点的值（用数组返回）。
func printListFromTailToHead(head *ListNode) []int {
	res := make([]int, 0)
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}
