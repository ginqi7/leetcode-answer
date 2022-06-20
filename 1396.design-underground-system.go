package main

type UndergroundSystem struct {
	stationNameMap map[int]string
	timeMap map[int]int
	sumTimeMap map[string]int
	countMap map[string]int
}


func Constructor() UndergroundSystem {
    return UndergroundSystem{make(map[int]string), make(map[int]int), make(map[string]int), make(map[string]int)}
}


func (this *UndergroundSystem) CheckIn(id int, stationName string, t int)  {
    this.stationNameMap[id] = stationName
	this.timeMap[id] = t
}


func (this *UndergroundSystem) CheckOut(id int, stationName string, t int)  {
	key := this.stationNameMap[id] + ":" + stationName
	this.sumTimeMap[key] += t - this.timeMap[id]
	this.countMap[key]++
}


func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
    key := startStation + ":" + endStation
	return float64(this.sumTimeMap[key]) / float64 (this.countMap[key])
}


/**
 * Your UndergroundSystem object will be instantiated and called as such:
 * obj := Constructor();
 * obj.CheckIn(id,stationName,t);
 * obj.CheckOut(id,stationName,t);
 * param_3 := obj.GetAverageTime(startStation,endStation);
 */
