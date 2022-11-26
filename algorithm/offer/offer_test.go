package offer

import (
	"fmt"
	"testing"
)

func Test_offer03(t *testing.T) {
	fmt.Println(findRepeatNumber([]int{2, 3, 1, 0, 2, 5, 3}))
}

func Test_offer4(t *testing.T) {
	fmt.Println(findNumberIn2DArray([][]int{
		{1, 2, 8, 9},
		{2, 4, 9, 12},
		{4, 7, 10, 13},
		{6, 8, 11, 15},
	}, 7))
	fmt.Println(findNumberIn2DArray([][]int{
		{1, 2, 8, 9},
		{2, 4, 9, 12},
		{4, 7, 10, 13},
		{6, 8, 11, 15},
	}, 3))
}

func Test_offer05(t *testing.T) {
	fmt.Println(replaceSpace("We Are Happy"))
	fmt.Println(replaceSpace("1 二 three xxx"))
}

func Test_offer06(t *testing.T) {
	fmt.Println(reversePrint(&ListNode{
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

func Test_offer7(t *testing.T) {
	x := buildTree([]int{1, 2, 4, 7, 3, 5, 6, 8}, []int{4, 7, 2, 1, 5, 3, 8, 6})
	fmt.Println(x)
}
func Test_offer10_1(t *testing.T) {
	fmt.Println(fib(2))
	fmt.Println(fib(5))
}
func Test_offer10_2(t *testing.T) {
	fmt.Println(numWays(2))
	fmt.Println(numWays(7))
	fmt.Println(numWays(0))
}
func Test_offer11(t *testing.T) {
	fmt.Println(minArray([]int{3, 4, 5, 1, 2})) // 1
	fmt.Println(minArray([]int{2, 2, 2, 0, 1})) // 2
}
func Test_offer12(t *testing.T) {
	fmt.Println(exist([][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}, "ABCCED")) // true
	fmt.Println(exist([][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}, "ABCB")) // false
	fmt.Println(exist([][]byte{
		{'a', 'b'},
		{'c', 'd'},
	}, "abcd")) // false
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
	fmt.Println(verifyPostorder([]int{1, 3, 2, 6, 5}))
	fmt.Println(verifyPostorder([]int{4, 8, 6, 12, 16, 14, 10}))
	fmt.Println(verifyPostorder([]int{1, 2, 5, 10, 6, 9, 4, 3}))
}
