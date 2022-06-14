package main
type OrderedStream struct {
    index int
	buffer []string
}


func Constructor(n int) OrderedStream {
    return OrderedStream{1, make([]string, n+1)}
}


func (this *OrderedStream) Insert(idKey int, value string) []string {
	ans := []string{}
	n := len(this.buffer)
    this.buffer[idKey] = value
	for this.index < n {
		val := this.buffer[this.index]
		if val != "" {
			ans = append(ans, val)
			this.index++
		} else {
			break
		}
	}
	return ans;
}


/**
 * Your OrderedStream object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Insert(idKey,value);
 */
