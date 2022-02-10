package leetcode

func reverse8(x int) int {
	tmp := 0
	for x != 0 {
		a := x % 10
		tmp = tmp * 10 + a
		x = x / 10
	}
	if tmp > 1<<31-1 || tmp < -(1<<31) {
		return 0
	}
	return tmp
}
