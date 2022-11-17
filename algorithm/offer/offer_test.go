package offer

import (
	"fmt"
	"testing"
)

func Test_offer01(t *testing.T) {
	fmt.Println(duplicate([]int{2, 3, 1, 0, 2, 5, 3}))
}

func Test_offer02(t *testing.T) {
	fmt.Println(Find(7, [][]int{
		{1, 2, 8, 9},
		{2, 4, 9, 12},
		{4, 7, 10, 13},
		{6, 8, 11, 15},
	}))
	fmt.Println(Find(3, [][]int{
		{1, 2, 8, 9},
		{2, 4, 9, 12},
		{4, 7, 10, 13},
		{6, 8, 11, 15},
	}))
}

func Test_offer03(t *testing.T) {
	fmt.Println(replaceSpace("We Are Happy"))
	fmt.Println(replaceSpace("1 二 three xxx"))
}

func Test_offer04(t *testing.T) {
	fmt.Println(printListFromTailToHead(&ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 5,
							Next: &ListNode{
								Val:  6,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}))
}

func Test_offer05(t *testing.T) {
	x := reConstructBinaryTree([]int{1, 2, 4, 7, 3, 5, 6, 8}, []int{4, 7, 2, 1, 5, 3, 8, 6})
	fmt.Println(x)
}

func Test_offer25(t *testing.T) {
	// 输入：1->2->4, 1->3->4
	// 输出：1->1->2->3->4->4
	fmt.Println(mergeTwoLists(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
			},
		},
	}, &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
			},
		},
	}))
}

func Test_offer33(t *testing.T) {
	// 输入: [1,6,3,2,5]
	// 输出: false
	fmt.Println(verifyPostorder([]int{1, 6, 3, 2, 5}))
	// 输入: [1,3,2,6,5]
	// 输出: true
	fmt.Println(verifyPostorder([]int{1,3,2,6,5}))
	fmt.Println(verifyPostorder([]int{4, 8, 6, 12, 16, 14, 10}))
	fmt.Println(verifyPostorder([]int{1,2,5,10,6,9,4,3}))
}
