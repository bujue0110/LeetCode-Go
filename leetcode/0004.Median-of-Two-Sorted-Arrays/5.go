package leetcode

func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	a := 0
	b := 0

	array := make([]int, 0)
	for a < m && b < n {
		if nums1[a] < nums2[b] {
			array = append(array, nums1[a])
			a++
		} else {
			array = append(array, nums2[b])
			b++
		}
	}
	for a < m {
		array = append(array, nums1[a])
		a++
	}
	for b < n {
		array = append(array, nums2[b])
		b++
	}

	if m+n%2 == 0 {
		a := (m+n)/2 - 1
		b := a
		return float64((a + b) / 2)
	} else {
		return float64(array[(m+n)/2])
	}

}
