package main
type MyQueue struct {
    s1 []int
	s2 []int
}


func Constructor() MyQueue {
	return MyQueue{[]int{}, []int{}}
}


func (this *MyQueue) Push(x int)  {
    this.s1 = append(this.s1, x)
}


func (this *MyQueue) Pop() int {
    ans := this.Peek()
	this.s2 = this.s2[:len(this.s2)-1]
	return ans
}


func (this *MyQueue) Peek() int {
    if len(this.s2) == 0 {
		for len(this.s1) != 0 {
			this.s2 = append(this.s2, this.s1[len(this.s1)-1])
			this.s1 = this.s1[:len(this.s1)-1]
		}
	}
	if len(this.s2) == 0 {
		return 0
	}
	ans := this.s2[len(this.s2)-1]
	return ans
}


func (this *MyQueue) Empty() bool {
    return len(this.s1) ==0 && len(this.s2) == 0
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
