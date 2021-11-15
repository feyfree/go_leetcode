package p0001

func twoSum(nums []int, target int) []int {
	var result []int
	for i := 0; i < len(nums); i++ {
		remaining := target - nums[i]
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == remaining {
				result = []int{i, j}
			}
		}
	}
	return result
}
