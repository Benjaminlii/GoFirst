package offer

// 剑指 Offer 3. 替换空格
// 请实现一个函数，将一个字符串s中的每个空格替换成“%20”。
// 例如，当字符串为We Are Happy.则经过替换之后的字符串为We%20Are%20Happy。
func replaceSpace(s string) string {
	res := []rune{}
	for _, r := range s {
		if r == rune(' ') {
			res = append(res, []rune("%20")...)
		} else {
			res = append(res, r)
		}
	}
	return string(res)
}
