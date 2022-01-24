package p0088

func merge(nums1 []int, m int, nums2 []int, n int) {
	end := m + n - 1
	m--
	n--
	for end >= 0 && m >= 0 && n >= 0 {
		if nums1[m] > nums2[n] {
			nums1[end] = nums1[m]
			m--
		} else {
			nums1[end] = nums2[n]
			n--
		}
		end--
	}
	for n >= 0 {
		nums1[end] = nums2[n]
		end--
		n--
	}
}
