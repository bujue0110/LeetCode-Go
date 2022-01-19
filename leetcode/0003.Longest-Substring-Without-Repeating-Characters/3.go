package leetcode

func lengthOfLongestSubstring3(s string) int {
	m := make(map[uint8]int)
	a := 0
	b := 0
	count := 0
	for a < len(s) {
		v := m[s[a]]
		if v > 0 {
			m[s[b]] = 0
			b++
		} else {
			m[s[a]] = 1
			a++
		}
		if a-b>count {
			count = a-b
		}
	}
	return count
}
