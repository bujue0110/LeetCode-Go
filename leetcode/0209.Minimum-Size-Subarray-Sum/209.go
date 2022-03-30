package leetcode

import "math"

// para209{7, []int{2, 3, 1, 2, 4, 3}},
func minSubArrayLen1(target int, nums []int) int {
	min := len(nums)
	for i, num := range nums {
		sum := num
		j := i + 1
		c := 1
		for sum < target && j <= len(nums)-1 {
			sum += nums[j]
			c++
			j++
		}
		if sum >= target {
			min = int(math.Min(float64(c), float64(min)))
		}
	}
	return min
}

// para209{7, []int{2, 3, 1, 2, 4, 3}},
// 2,5,6,8,12,15
func minSubArrayLen2(target int, nums []int) int {
	min := len(nums)
	nums2 := make([]int, 0)
	nums2[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		nums2[i] = nums[i]+nums2[i-1]
	}
	return min
}
