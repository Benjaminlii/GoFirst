package offer

// 剑指 Offer 2. 二维数组中的查找
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
