package p0080

func removeDuplicates(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	} else {
		// 相当于新的访问索引
		i := 0
		// 重复数量
		repeat := 2
		// 当前重复值
		currentRepeat := 1
		for j := 1; j < length; j++ {
			if nums[j] == nums[i] && currentRepeat < repeat {
				i++
				nums[i] = nums[j]
				currentRepeat++
			} else if nums[j] != nums[i] {
				if currentRepeat >= repeat {
					currentRepeat = 1
				}
				i++
				nums[i] = nums[j]
			}
		}
		return i + 1
	}
}
