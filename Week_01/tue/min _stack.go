package tue

import "math"

// lc 155 最小栈
type MinStack struct {
	stack []int
	min_s []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack: []int{},
		min_s: []int{math.MaxInt64},
	}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	tmp := this.min_s[len(this.min_s)-1]
	this.min_s = append(this.min_s, min(int(tmp), x))
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.min_s = this.min_s[:len(this.min_s)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

// 就是比普通stacks 多了衣蛾常数时间复杂度检索出最小元素的api
func (this *MinStack) GetMin() int {
	return this.min_s[len(this.min_s)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
