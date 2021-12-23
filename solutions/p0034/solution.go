package p0034

func searchRange(nums []int, target int) []int {
	n := len(nums)
	if n == 0 {
		return []int{-1, -1}
	}
	return []int{firstPosition(nums, target), lastPosition(nums, target)}
}

func firstPosition(nums []int, target int) int {
	low := 0
	high := len(nums)
	for low < high {
		mid := low + (high-low)/2
		if nums[mid] >= target {
			high = mid
		} else {
			low = mid + 1
		}
	}
	if high == len(nums) || nums[high] != target {
		return -1
	}
	return high
}

func lastPosition(nums []int, target int) int {
	low := 0
	high := len(nums)
	for low < high {
		mid := low + (high-low)/2
		if nums[mid] > target {
			high = mid
		} else {
			low = mid + 1
		}
	}
	low--
	if low < 0 || nums[low] != target {
		return -1
	}
	return low
}
