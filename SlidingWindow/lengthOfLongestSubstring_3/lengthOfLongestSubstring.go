package lengthOfLongestSubstring_3

// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。

// hot 100 No.3 无重复字符的最长子串
func lengthOfLongestSubstring(s string) int {
	maxSubStrLen := 0              // 当前找到的无重复字符的最长子串的长度
	subStrCharMp := map[rune]int{} // 子串中包含的字符到其下标的映射

	l, r := 0, 0 // 滑动窗口的左右边界
	for r < len(s) {
		// 如果右边新加入的是当前子串中没重复的字符
		if _, exist := subStrCharMp[rune(s[r])]; !exist {
			subStrCharMp[rune(s[r])] = r
			r++
			curLen := len(subStrCharMp)
			if curLen > maxSubStrLen {
				maxSubStrLen = curLen
			}
			continue
		}

		// 如果右边新加入的字符在当前子串中已经存在
		// 将l右移到子串中该字符的下一个字符位置，并在map中删除这之间的所有字符的映射
		for l <= subStrCharMp[rune(s[r])] {
			delete(subStrCharMp, rune(s[l]))
			l++
		}
		subStrCharMp[rune(s[r])] = r
		r++
	}

	return maxSubStrLen
}
