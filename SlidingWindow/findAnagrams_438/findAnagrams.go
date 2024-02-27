package findAnagrams_438

// 给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
// 异位词 指由相同字母重排列形成的字符串（包括相同的字符串）。

// hot 100 No.438 找到字符串中所有字母异位词
// [方法三] 滑动窗口，用数组来记录滑动窗口内与p的字符出现的次数差值，然后用变量diff来记录它们出现次数不相等的字符的个数
// diff为0时代表成功找到异位词
func findAnagrams(s string, p string) []int {
	var res []int
	if len(s) < len(p) {
		return res
	}

	var count [26]int // 用数组来记录滑动窗口内子串与p的字符出现的次数差值
	for i, ch := range p {
		count[s[i]-'a']++
		count[ch-'a']--
	}

	diff := 0 // 记录滑动窗口内子串与p的出现次数不相等的字符的个数
	for _, n := range count {
		if n != 0 {
			diff++
		}
	}

	if diff == 0 {
		res = append(res, 0)
	}

	for i, ch := range s[:len(s)-len(p)] {
		if count[ch-'a'] == 1 { // 如果要离开的本来就是多的那个，则离开后该字符的出现次数就相等了
			diff--
		}
		if count[ch-'a'] == 0 { // 如果要离开的字符本来就是次数相等的，则离开后反而不等了
			diff++
		}
		count[ch-'a']-- // 左边将要离开滑动窗口的字符

		if count[s[i+len(p)]-'a'] == -1 {
			diff--
		}
		if count[s[i+len(p)]-'a'] == 0 {
			diff++
		}
		count[s[i+len(p)]-'a']++ // 右边新加入的字符

		if diff == 0 {
			res = append(res, i+1)
		}
	}

	return res
}

//// [方法二] 滑动窗口，用数组来表示字符出现的次数（题目写了仅包含小写字母，因此用长度为26的数组即可）
//// 另外，在 Go 中，可以使用==来比较数组是否相等，只有当两个数组的每个元素都相等时，它们才会相等。这种比较方式适用于固定长度的数组。
//// 如果是切片（slice），则不能直接使用 == 进行比较，需要逐个元素进行比较。
//func findAnagrams(s string, p string) []int {
//	var res []int
//	if len(s) < len(p) {
//		return res
//	}
//
//	var sCnt, pCnt [26]int // 技巧：用数组来记录字符出现的次数
//	for i, ch := range p {
//		pCnt[ch-'a']++ // 字符转化为数组下标
//		sCnt[s[i]-'a']++
//	}
//
//	// 固定长度的数组能够用==直接比较
//	if pCnt == sCnt {
//		res = append(res, 0)
//	}
//
//	// 窗口起始为i，长度固定，就能唯一确定一个滑动窗口
//	for i, ch := range s[:len(s)-len(p)] {
//		// 窗口每次滑动一格，因此只用更新sCnt中新增和去除的两个字符的出现次数，大大减少了每次都重新统计一遍子串的字符出现次数！
//		sCnt[ch-'a']--
//		sCnt[s[i+len(p)]-'a']++
//		if pCnt == sCnt {
//			res = append(res, i+1) // 注意此次循环的滑动窗口起始位置其实是i+1
//		}
//	}
//
//	return res
//}

//// [方法一] 滑动窗口，时间和空间复杂度很大，推测是循环太多导致的
//func findAnagrams(s string, p string) []int {
//	var res []int
//	if len(s) < len(p) {
//		return res
//	}
//
//	pMp := map[rune]int{} // p中字符到出现次数的映射
//	for _, char := range p {
//		pMp[char]++
//	}
//
//	l, r := 0, len(p)-1
//	for r < len(s) {
//		// 检查该窗口内的字符是否与p中的字符及对应的出现次数都相同
//		// 从右往左检查，如果遇到p中不存在的字符则不用继续检查了，直接将左边界移到该字符的下一个字符
//		mp := map[rune]int{} // 记录此滑动窗口中的字符出现次数
//		i := r
//		for i >= l {
//			if _, exist := pMp[rune(s[i])]; !exist {
//				break
//			}
//			mp[rune(s[i])]++
//			i--
//		}
//
//		// 若是遇到p中未出现的字符
//		if i >= l {
//			l = i + 1
//			r = l + len(p) - 1
//			continue
//		}
//
//		if len(mp) != len(pMp) {
//			l++
//			r++
//			continue
//		}
//
//		ok := true
//		// 若没有p中未出现的字符，那么还要检查各字符得出现次数是否一致
//		for char, num := range pMp {
//			if mp[char] != num { // 只要有次数不一致的，窗口右滑
//				ok = false
//				break
//			}
//		}
//
//		if ok { // 满足异位词，记录子串的起始索引
//			res = append(res, l)
//		}
//		l++
//		r++
//	}
//
//	return res
//}
