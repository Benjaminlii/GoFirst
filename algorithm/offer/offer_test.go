package offer

import (
	"fmt"
	"testing"
)

func Test_duplicate(t *testing.T) {
	fmt.Println(duplicate([]int{2, 3, 1, 0, 2, 5, 3}))
}

func TestFind(t *testing.T) {
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

func Test_replaceSpace(t *testing.T) {
	fmt.Println(replaceSpace("We Are Happy"))
	fmt.Println(replaceSpace("1 二 three xxx"))
}

func Test_printListFromTailToHead(t *testing.T) {
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


func Test_reConstructBinaryTree(t *testing.T) {
	x := reConstructBinaryTree([]int{1, 2, 4, 7, 3, 5, 6, 8}, []int{4, 7, 2, 1, 5, 3, 8, 6})
	fmt.Println(x)
}
