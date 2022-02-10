package leetcode

func isPalindrome2(x int) bool {
	if x <= 10 {
		return false
	}
	array := make([]int, 0, 32)
	if x != 0 {
		array = append(array, x %10)
		x =x /10
	}
	i := 0
	j := len(array)

	for i<j {
		if array[i] != array[j] {
			return false
		}
		i++
		j--
	}
	return true
}
