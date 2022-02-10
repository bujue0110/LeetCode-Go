package leetcode

import "math"

func maxArea1(height []int) int {
	i := 0
	j := len(height) - 1
	var max float64
	for i < j {
		width := j-i
		if height[i] < height[j] {
			max = math.Max(float64(height[i]*width), max)
			i++
		} else {
			max = math.Max(float64(height[j]*width), max)
			j--
		}
	}
	return int(max)
}
