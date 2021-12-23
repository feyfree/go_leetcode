package p0035

func searchInsert(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	// 二分查找 first item > target 当前的index
	start := 0
	end := n - 1
	for start <= end {
		mid := start + (end-start)/2
		if nums[mid] >= target {
			if mid-1 >= 0 && nums[mid-1] < target {
				return mid
			}
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return end + 1
}
