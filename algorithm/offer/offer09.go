package offer

// 剑指 Offer 09. 用两个栈实现队列
// 用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，
// 分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1 )

type CQueue struct {
	InputStack  []int
	OutputStack []int
}

func Constructor() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	this.InputStack = append(this.InputStack, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.OutputStack) == 0 {
		for i := len(this.InputStack) - 1; i >= 0; i-- {
			this.OutputStack = append(this.OutputStack, this.InputStack[i])
		}
		this.InputStack = make([]int, 0)
	}
	if len(this.OutputStack) == 0 {
		return -1
	}
	res := this.OutputStack[len(this.OutputStack)-1]
	this.OutputStack = this.OutputStack[:len(this.OutputStack)-1]
	return res
}
