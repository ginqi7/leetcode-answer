package main
func moveZeroes(nums []int)  {
    n := len(nums)
	if n == 0 {
		return
	}
	count := 0
	for i:= 0; i < n; i++ {
		if nums[i] != 0 {
			nums[count] = nums[i]
			count++
		}
	}
	
	for j := count; j < n; j++ {
		nums[j] = 0
	} 
}
