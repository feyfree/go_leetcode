package p0033

func search(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	start := 0
	end := n - 1
	for start <= end {
		mid := start + (end-start)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < nums[end] {
			if target > nums[mid] && target <= nums[end] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		} else if nums[mid] > nums[end] {
			if target >= nums[start] && target < nums[mid] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else {
			end--
		}
	}
	return -1
}
