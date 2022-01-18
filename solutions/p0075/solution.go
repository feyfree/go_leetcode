package p0075

func sortColors(nums []int) {
	counter := make([]int, 3)
	for _, num := range nums {
		counter[num]++
	}
	index := 0
	for i := 0; i < 3; i++ {
		for counter[i] > 0 {
			nums[index] = i
			index++
			counter[i]--
		}
	}
}
