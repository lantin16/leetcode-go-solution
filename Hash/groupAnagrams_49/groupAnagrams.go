package groupAnagrams_49

import "sort"

// 给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
// 字母异位词 是由重新排列源单词的所有字母得到的一个新单词。

// hot 100 No.49 字母异位词分组
func groupAnagrams(strs []string) [][]string {
	//mp := map[[26]int][]string{}	// 使用数组作为哈希表的键
	var res [][]string
	groupMap := make(map[string][]string) // 值为切片的map，{排序好的string, 对应单词slice}
	for _, str := range strs {
		anagramStr := sortChar1(str)
		//anagramStr := sortChar2(str)
		if wordList, ok := groupMap[anagramStr]; ok { // 如果该异位词分组已存在
			wordList = append(wordList, str) // 追加进去新词到对应list
			groupMap[anagramStr] = wordList
		} else { // 如果该异位词分组不存在则新建分组
			l := []string{str}
			groupMap[anagramStr] = l
		}
	}

	for _, list := range groupMap {
		res = append(res, list)
	}

	return res
}

// 将string中的字母重排序，按照a-z从小到大，手动冒泡
func sortChar1(str string) string {
	charArr := []byte(str)
	// 冒泡排序
	for i := 0; i < len(charArr); i++ {
		for j := len(charArr) - 1; j > i; j-- {
			if charArr[j] < charArr[j-1] {
				charArr[j-1], charArr[j] = charArr[j], charArr[j-1] // go里可以直接交换
			}
		}
	}
	res := string(charArr)
	return res
}

// 定义序列类型
type byteSlice []byte

// 实现三要素方法：Len()，Less()，Swap()
func (b byteSlice) Len() int {
	return len(b)
}

func (b byteSlice) Less(i, j int) bool {
	return b[i] < b[j]
}

func (b byteSlice) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

// 将string中的字母重排序，按照a-z从小到大，实现三要素方法调用sort
func sortChar2(str string) string {
	charArr := []byte(str)
	sort.Sort(byteSlice(charArr))
	res := string(charArr)
	return res

}
