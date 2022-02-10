package leetcode

func longestPalindrome4(s string) string {
	res := ""
	for i := range s {
		res = max(s, i, i, res)
		res = max(s, i, i+1, res)
	}
	return res
}

func max(s string, i,j int, res string) string {
	sub := ""
	for i>=0 && j <= len(s) -1 && s[i]==s[j]{
		sub = s[i:j+1]
		i--
		j++
	}
	if len(sub) > len(res) {
		return sub
	}
	return res
}
