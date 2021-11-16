package p0004

import "math"

// 变式的二分搜索
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 方便处理我们可以确保nums1 的长度始终小于 nums2 的长度
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}
	mid := (len(nums1) + len(nums2) + 1) / 2
	l := 0
	r := len(nums1)
	for l < r {
		m1 := l + (r-l)/2
		m2 := mid - m1
		if nums1[m1] < nums2[m2-1] {
			l = m1 + 1
		} else {
			r = m1
		}
	}
	// m1 和 m2 是两个数组里面能够参与计算的最高位置(从1 开始)
	m1 := l
	m2 := mid - l

	// 如果是奇数个数组的话, 只需要参与计算的最大位置(注意是m1 - 1)
	p1a := math.MinInt32
	if m1 > 0 {
		p1a = nums1[m1-1]
	}
	p2a := math.MinInt32
	if m2 > 0 {
		p2a = nums2[m2-1]
	}

	a := float64(int(math.Max(float64(p1a), float64(p2a))))

	if (len(nums2)+len(nums1))&1 == 1 {
		return a
	}

	// 如果是偶数个数组的话, 后推一下取最小值
	p1b := math.MaxInt32
	if m1 < len(nums1) {
		p1b = nums1[m1]
	}
	p2b := math.MaxInt32
	if m2 < len(nums2) {
		p2b = nums2[m2]
	}

	b := math.Min(float64(p1b), float64(p2b))
	return (a + b) * 0.5

}
