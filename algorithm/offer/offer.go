package offer

// 1. 数组中重复的数字
// 在一个长度为n的数组里的所有数字都在0到n-1的范围内。
// 数组中某些数字是重复的，但不知道有几个数字是重复的。也不知道每个数字重复几次。
// 请找出数组中任意一个重复的数字。
// 例如，如果输入长度为7的数组[2,3,1,0,2,5,3]，那么对应的输出是2或者3。存在不合法的输入的话输出-1
func duplicate(numbers []int) int {
	mark := make([]int, len(numbers))
	for _, v := range numbers {
		if mark[v] > 0 {
			return v
		}
		mark[v]++
	}
	return -1
}

// 2. 二维数组中的查找
// 在一个二维数组array中（每个一维数组的长度相同），每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。
// 请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
// [
//
//	[1,2,8,9],
//	[2,4,9,12],
//	[4,7,10,13],
//	[6,8,11,15]
//
// ]
// 给定 target = 7，返回 true。
// 给定 target = 3，返回 false
// 进阶：空间复杂度 O(1) ，时间复杂度 O(n+m)
func Find(target int, array [][]int) bool {
	res := false
	for i, j := len(array)-1, 0; i >= 0 && j < len(array[0]); {
		getNum := array[i][j]
		if getNum == target {
			res = true
			break
		} else if getNum > target {
			i--
		} else {
			j++
		}
	}
	return res
}

// 3. 替换空格
// 请实现一个函数，将一个字符串s中的每个空格替换成“%20”。
// 例如，当字符串为We Are Happy.则经过替换之后的字符串为We%20Are%20Happy。
func replaceSpace(s string) string {
	res := []rune{}
	for _, r := range s {
		if r == rune(' ') {
			res = append(res, []rune("%20")...)
		} else {
			res = append(res, r)
		}
	}
	return string(res)
}

// 4. 从尾到头打印链表
// 输入一个链表的头节点，按链表从尾到头的顺序返回每个节点的值（用数组返回）。
type ListNode struct {
	Val  int
	Next *ListNode
}

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

// 5. 重建二叉树
// 给定节点数为 n 的二叉树的前序遍历和中序遍历结果，请重建出该二叉树并返回它的头结点。
// 例如输入前序遍历序列{1,2,4,7,3,5,6,8}和中序遍历序列{4,7,2,1,5,3,8,6}，则重建出如下图所示。
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func reConstructBinaryTree(pre []int, vin []int) *TreeNode {
	if len(pre)+len(vin) == 0 {
		return nil
	}

	root := &TreeNode{
		Val: pre[0],
	}

	var leftVin []int
	var rightVin []int
	var leftPre []int
	var rightPre []int
	for i, v := range vin {
		if v == root.Val {
			leftVin = vin[0:i]
			rightVin = vin[i+1:]
			leftPre = pre[1 : len(leftVin)+1]
			rightPre = pre[1+len(leftPre):]
			break
		}
	}
	root.Left = reConstructBinaryTree(leftPre, leftVin)
	root.Right = reConstructBinaryTree(rightPre, rightVin)
	return root
}
