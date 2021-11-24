package p0026

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	} else {
		i := 0
		for j := 1; j < n; j++ {
			if nums[j] != nums[i] {
				i++
				nums[i] = nums[j]
			}
		}
		return i + 1
	}
}
