package p0041

//这道题可以认为是一个pattern 和 一个数组num 进行匹配的一个结果, 如果nums 的长度为n的话
//
//pattern: [1, 2, 3, 4, ... n]
//
//我们将nums 这个数组 往这个pattern 上面靠， 就是nums对应的索引尽可能放pattern上面对应索引的值，对nums原地修改如果能匹配的上就放。
func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			// 假设 k = nums[i] ; 进行交换 k 应该出现在 nums[k - 1] 上面
			temp := nums[nums[i]-1]
			nums[nums[i]-1] = nums[i]
			nums[i] = temp
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}
