package offer

// 剑指 Offer 5. 重建二叉树
// 给定节点数为 n 的二叉树的前序遍历和中序遍历结果，请重建出该二叉树并返回它的头结点。
// 例如输入前序遍历序列{1,2,4,7,3,5,6,8}和中序遍历序列{4,7,2,1,5,3,8,6}，则重建出如下图所示。
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
