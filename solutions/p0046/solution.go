package p0046

func permute(nums []int) [][]int {
	var result [][]int
	used := make([]bool, len(nums))
	var current []int
	permute0(nums, 0, len(nums), &used, current, &result)
	return result
}

func permute0(nums []int, d int, n int, used *[]bool, current []int, result *[][]int) {
	if d == n {
		*result = append(*result, append([]int(nil), current...))
		return
	}
	for i := 0; i < n; i++ {
		if (*used)[i] {
			continue
		}
		(*used)[i] = true
		permute0(nums, d+1, n, used, append(current, nums[i]), result)
		(*used)[i] = false
	}
}
