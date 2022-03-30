package leetcode

import "sort"

func searchRange1(nums []int, target int) []int {
	l := sort.SearchInts(nums, target)
	r := sort.SearchInts(nums, target +1) -1
	return []int{l, r}
}

