package p0031

func nextPermutation(nums []int) {
	n := len(nums)
	if n == 0 {
		return
	}
	i := n - 2
	for i >= 0 && nums[i+1] <= nums[i] {
		i--
	}
	if i >= 0 {
		j := n - 1
		for j >= 0 && nums[j] <= nums[i] {
			j--
		}
		swap(nums, i, j)
	}
	reverse(nums, i+1, n-1)
}

func swap(nums []int, i int, j int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}

func reverse(nums []int, start int, end int) {
	for start < end {
		swap(nums, start, end)
		start++
		end--
	}
}
